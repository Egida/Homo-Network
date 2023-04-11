package server

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
)

func Reboot() {

	defer func() { // try catch
		if er := recover(); er != nil {
			println("Reboot: ")
			fmt.Print(er)
			return
		}
	}()

	data, err := os.ReadFile("./data/servers.txt")
	if err != nil {
		fmt.Println("Read servers: " + err.Error())
	}

	for _, i := range strings.Split(string(data), "\n") {
		serv := strings.Split(string(i), ":")

		if len(serv) < 2 {
			continue
		}

		wg.Add(1)
		go func() {
			ses, err := sshNew(serv[1], serv[2], serv[0])
			if err != nil {
				fmt.Println("[HOMO NET] Can't reboot: " + serv[0] + ":22")
				wg.Done()
				return
			}

			ses.reboot(&wg)
		}()
		wg.Wait()
	}
}

func (sess *Sshsess) reboot(w *sync.WaitGroup) {

	sshSesh, _ := sess.Session.NewSession()
	var setSession bytes.Buffer
	sshSesh.Stdout = &setSession

	fmt.Println("[HOMO NET] Rebooted: " + sess.Ip + ":22")

	go func() {
		sshSesh.Run("reboot")
		sshSesh.Close()

	}()
	wg.Done()
}
