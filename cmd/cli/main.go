package main

import (
	"flag"
	mathClient "github.com/pojntfx/rpcx-math/math/client"
	"github.com/smallnest/rpcx/client"
	"log"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	discovery := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Math", client.Failtry, client.RandomSelect, discovery, client.DefaultOption)
	defer xclient.Close()

	addedNumbers, _ := mathClient.Add(xclient, 1, 3)
	subtractedNumbers, _ := mathClient.Subtract(xclient, 1, 3)

	log.Println("Added numbers:", addedNumbers)
	log.Println("Subtracted numbers:", subtractedNumbers)
}
