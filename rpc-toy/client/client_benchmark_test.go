package main

import (
	"goURL-shortie/rpc-toy/common"
	"net/rpc"
	"testing"
)

// I am benchmarking the calls and not the setup.

func BenchmarkRpcClientCall(b *testing.B) {
	// setup code - dial RPC server
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatalf("Failed to dial RPC server: %v", err)
	}
	defer client.Close()

	args := &common.Args{A: 12, B: 89}
	var reply common.Reply
	b.ResetTimer() // reset timer to exclude setup time

	for i := 0; i < b.N; i++ {
		err := client.Call("Calculator.Add", args, &reply)
		if err != nil {
			b.Fatalf("RPC call failed: %v", err)
		}

	}
}

// **** --------Adding RPC payload knob (tiny vs 64KB)---------------------- ****
func BenchmarkRpcClientCall_TinyPayload(b *testing.B) {
	// setup code - dial RPC server
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatal("Failed to dial RPC server: ", err)
	}
	defer client.Close()
	args := &common.Args{A: 10, B: 20, Payload: make([]byte, 16)}
	var reply common.Reply
	b.ReportAllocs()
	expected := 10 + 20 + 16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := client.Call("Calculator.Add", args, &reply)
		if err != nil {
			b.Fatal("RPC call failed: ", err)
		}
		if reply.Result != expected {
			b.Errorf("Unexpected reply: got %d, want %d", reply.Result, expected)
		}
	}

}

func BenchmarkRpcClientCall_LargePayload(b *testing.B) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatal("Failed to dial RPC server: ", err)
	}
	defer client.Close()
	buf := make([]byte, 64*1024)
	args := &common.Args{A: 10, B: 20, Payload: buf}
	var reply common.Reply
	b.ReportAllocs()
	expected := 10 + 20 + 65536
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := client.Call("Calculator.Add", args, &reply)
		if reply.Result != expected {
			b.Errorf("Unexpected reply: got %d, want %d", reply.Result, expected)
		}
		if err != nil {
			b.Fatal("RPC call failed: ", err)
		}
	}
}

// **** -- Large Payload Parallel Benchmarking -- ****

func BenchmarkRpcClientCall_LargePayload_Parallel(b *testing.B) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatal("Failed to dialRPC server", err)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buf := make([]byte, 64*1024)
			args := &common.Args{A: 10, B: 20, Payload: buf}
			var reply common.Reply
			expected := 10 + 20 + 65536
			err := client.Call("Calculator.Add", args, &reply)
			if reply.Result != expected {
				b.Errorf("Unexpected reply: got %d, want %d", reply.Result, expected)
			}
			if err != nil {
				b.Fatal("RPC call failed: ", err)
			}
		}
	})
}

// Large Payload Allocate Per call
func BenchmarkRpcClientCall_LargePayload_AllocPerCall(b *testing.B) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatal("Failed to dial RPC server: ", err)
	}
	b.ResetTimer()
	var reply common.Reply
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 64*1024)
		args := &common.Args{A: 10, B: 20, Payload: buf}
		b.ReportAllocs()
		expected := 10 + 20 + 65536
		err := client.Call("Calculator.Add", args, &reply)
		if reply.Result != expected {
			b.Errorf("Unexpected reply: got %d, want %d", reply.Result, expected)
		}
		if err != nil {
			b.Fatal("RPC call failed: ", err)
		}

	}
}

// *** --- Benchmarking using shared Pool of Payload Buffers ---- ****

func BenchmarkRpcClientCall_LargePayload_BufferPool(b *testing.B) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatal("Failed to dial RPC server: ", err)
	}
	defer client.Close()
	pool := make(chan []byte, 100) // buffer pool channel

	// Pre-fill the pool with buffers
	for i := 0; i < 100; i++ {
		pool <- make([]byte, 64*1024)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {

		args := &common.Args{A: 10, B: 20}
		var reply common.Reply
		for pb.Next() {
			buf := <-pool
			args.Payload = buf
			expected := 10 + 20 + 65536
			err := client.Call("Calculator.Add", args, &reply)
			if reply.Result != expected {
				b.Errorf("Unexpected reply: got %d, want %d", reply.Result, expected)
			}
			if err != nil {
				b.Fatal("RPC call failed: ", err)
			}
			pool <- buf
		}
	})
}

// **** ---- Local benchmark without RPC overhead ---- ****

type Calculator struct {
}

func (c *Calculator) Add(args *common.Args, reply *common.Reply) error {
	reply.Result = args.A + args.B
	return nil
}

func BenchmarkLocalAdd(b *testing.B) {
	calculator := new(Calculator) // local instance of the service
	args := &common.Args{A: 12, B: 89}
	var reply common.Reply

	b.ResetTimer() // reset timer to exclude setup time

	for i := 0; i < b.N; i++ {
		err := calculator.Add(args, &reply)
		if err != nil {
			b.Fatalf("Local Add call failed: %v", err)
		}
	}
}

// Benchmark RPC with variable payload sizes
func BenchmarkRpcClientCall_VarPayloadSize(b *testing.B) {
	b.Run("16KB", func(b *testing.B) {
		Bechmark_with_provided_payloadSize(b, 16*1024)
	})
	b.Run("32KB", func(b *testing.B) {
		Bechmark_with_provided_payloadSize(b, 32*1024)
	})
	b.Run("64KB", func(b *testing.B) {
		Bechmark_with_provided_payloadSize(b, 64*1024)
	})
	b.Run("256KB", func(b *testing.B) {
		Bechmark_with_provided_payloadSize(b, 256*1024)
	})
	b.Run("1024KB", func(b *testing.B) {
		Bechmark_with_provided_payloadSize(b, 1024*1024)
	})
}

func Bechmark_with_provided_payloadSize(b *testing.B, payloadSize int) {
	client, err := rpc.Dial("tcp", "localhost:8000")
	if err != nil {
		b.Fatal("Failed to dial RPC server: ", err)
	}
	defer client.Close()
	buf := make([]byte, payloadSize)
	args := &common.Args{A: 10, B: 20, Payload: buf}
	var reply common.Reply
	b.ReportAllocs()
	expected := 10 + 20 + payloadSize
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := client.Call("Calculator.Add", args, &reply)
		if reply.Result != expected {
			b.Fatalf("Unexpected reply: got %d, want %d", reply.Result, expected)
		}
		if err != nil {
			b.Fatal("RPC call failed: ", err)
		}
	}
}
