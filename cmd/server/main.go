package main

import (
	"flag"
	"github.com/pojntfx/rpcx-math/math/service"
	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Math", new(service.Math), "")
	s.Serve("tcp", *addr)
}
