package main

import (
	"fmt"
	"log"

	"github.com/midknight24/netvis/internal/backend"
	"github.com/midknight24/netvis/internal/core"
)

func main() {
	interfaces := core.ScanInterface()
	fmt.Println("Available interfaces: ")
	for i, intf := range interfaces {
		fmt.Printf("%v) %v\n", i, intf.Description)
	}
	var intf int
	fmt.Println("Choose interface to capture: ")
	fmt.Scan(&intf)
	if intf < 0 || intf >= len(interfaces) {
		log.Fatal("Invalid interface")
	}
	server := backend.New(interfaces[intf].Name)
	server.Run()
}
