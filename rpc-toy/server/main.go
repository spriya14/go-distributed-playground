package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

// This has no fields, it exists so we can attach methods to it.
// Calculator is the service, and its methods will be callable remotely.
type Calculator struct {
	addCalls      int
	multiplyCalls int
}

type Args struct {
	A, B int
}

type Reply struct {
	Result int
}

// exported method name must start with capital letter, it means it can be accessed from other packages.
// RPC method arguments and reply MUST be pointers.
// Also, RPC method must return an error type.
func (c *Calculator) Add(args *Args, reply *Reply) error {
	reply.Result = args.A + args.B
	c.addCalls++
	fmt.Println("Add method called", c.addCalls, "times")
	return nil
}

func (c *Calculator) Multiply(args *Args, reply *Reply) error {
	time.Sleep(5 * time.Second)
	reply.Result = args.A * args.B
	c.multiplyCalls++
	fmt.Println("Multiply method called", c.multiplyCalls, "times")
	return nil
}

func main() {
	calculator := new(Calculator) // new service instance
	rpc.Register(calculator)      // this means expose the methods of Calculator type as RPC methods.
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error starting RPC server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("RPC server listening on port 8000")
	rpc.Accept(listener) // Accept incoming connections and serve requests

}
