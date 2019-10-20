package main

import (
	"flag"
	"github.com/pojntfx/rpcx-math/math/clientlib"
	"github.com/smallnest/rpcx/client"
	"log"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	discovery := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Math", client.Failtry, client.RandomSelect, discovery, client.DefaultOption)
	defer xclient.Close()

	addedNumbers, addedNumbersErr := clientlib.Add(xclient, 1, 3)
	if addedNumbersErr != nil {
		log.Fatalln("Could not add numbers:", addedNumbersErr)
	} else {
		log.Println("Added numbers:", addedNumbers)
	}

	subtractedNumbers, subtractedNumbersErr := clientlib.Subtract(xclient, 1, 3)
	if subtractedNumbersErr != nil {
		log.Fatalln("Could not subtract numbers:", subtractedNumbersErr)
	} else {
		log.Println("Subtracted numbers:", subtractedNumbers)
	}

}
