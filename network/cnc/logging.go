package cnc

import (
	"homo/network/config"
	"io/ioutil"
	"net/http"
	"strings"
)

func Log(text string) {
	if !config.GetConfig().Logging.Logging {
		return
	}
	text = strings.ReplaceAll(text, "||", "%0D%0A")
	text = strings.ReplaceAll(text, ".", "\\.")
	ed, _ := http.Get("https://api.telegram.org/bot" + config.GetConfig().Logging.BotToken + "/sendMessage?chat_id=" + config.GetConfig().Logging.ChatId + "&parse_mode=MarkdownV2&text=" + text + "*")

	ioutil.ReadAll(ed.Body)

}
