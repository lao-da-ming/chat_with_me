package main

import "chat_with_me/app"

func main() {
	engine := app.InitApp()
	if err := app.Run(engine); err != nil {
		panic(err)
	}
}
