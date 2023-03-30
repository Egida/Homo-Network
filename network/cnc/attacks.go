package cnc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type attacksStruct struct {
	Duration int
	Finish   int
	Method   string
	Target   string
	Login    string
	Port     string
}

var AttacksMap = make(map[string]attacksStruct)

func NewAttack(login string, duration string, method string, target string, port string) {

	duration = strings.ReplaceAll(duration, "\x00", "")
	duration = strings.ReplaceAll(duration, "\n", "")
	duration = strings.ReplaceAll(duration, "\r", "")
	duration = strings.ReplaceAll(duration, "\x03", "")

	dur, err := strconv.Atoi(duration)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i <= dur; i++ {

		var str = attacksStruct{
			Duration: dur,
			Finish:   dur - i,
			Login:    login,
			Method:   method,
			Target:   target,
			Port:     port,
		}

		AttacksMap[login] = str
		if i == dur {
			delete(AttacksMap, str.Login)
		}

		time.Sleep(1 * time.Second)
	}

}
