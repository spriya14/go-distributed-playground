package main

import (
	"fmt"
	"goURL-shortie/rpc-toy/common"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8000") // think of it like db connection, open it once and reuse it, close it when done.
	if err != nil {
		fmt.Println("Error Connecting: ", err)
		return
	}
	args := &common.Args{A: 12, B: 89}
	var reply common.Reply
	var multiplicationReply common.Reply
	// RPC is not like normal function call - its serialization + transport disguised as a function call.
	// hence we need to pass args and reply as pointers.
	fmt.Println("Time Stamp Before Addition:", time.Now())
	err = client.Call("Calculator.Add", args, &reply) // & reply is the place to store the result
	fmt.Println("Time Stamp After Addition:", time.Now())

	fmt.Println("Time Stamp Before Multiplication:", time.Now())
	err = client.Call("Calculator.Multiply", args, &multiplicationReply)
	fmt.Println("Time Stamp After Multiplication:", time.Now())
	if err != nil {
		fmt.Println("Error calling RPC method: ", err)
		return
	}

	fmt.Printf("Result of %d + %d = %d\n", args.A, args.B, reply.Result)
	fmt.Printf("Result of %d * %d = %d\n", args.A, args.B, multiplicationReply.Result)

}
