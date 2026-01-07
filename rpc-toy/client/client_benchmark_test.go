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
