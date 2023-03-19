package main

import (
	"fmt"
	"homo/network/cnc"
	"homo/network/config"
	"homo/network/server"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	config := config.GetConfig()
	go server.StartServer(config)
	go cnc.Start(config)
	fmt.Println("Started")
	wg.Wait()
}
