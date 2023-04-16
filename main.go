package main

import (
	"fmt"
	"homo/network/cnc"
	"homo/network/config"
	"homo/network/server"
	"homo/network/server/api"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	if runtime.GOOS != "linux" {
		fmt.Println("Unsupported OS")
		return
	}

	wg.Add(1)
	config := config.GetConfig()
	go server.StartServer(config)
	go cnc.Start(config)
	go api.StartApi(config)
	wg.Wait()
}
