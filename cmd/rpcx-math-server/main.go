package main

import (
	"flag"
	"github.com/pojntfx/rpcx-math/math/svc"
	"github.com/smallnest/rpcx/server"
	"log"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Math", new(svc.Math), "")
	err := s.Serve("tcp", *addr)

	if err != nil {
		log.Fatalln("Could not start server:", err)
	}
}
