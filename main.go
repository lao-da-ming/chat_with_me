package main

import (
	"chat_with_me/app/web/cmd"
	"github.com/gin-gonic/gin"
)

func main() {
	app := cmd.WireApp(gin.Default())
	if err := app.Run(); err != nil {
		panic(err)
	}
}
