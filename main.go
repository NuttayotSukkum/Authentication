package main

import (
	"Authentication/Configs"
	"Authentication/Routers"
)

func main() {
	e := Routers.UserRouter()
	Configs.InitEnv()
	Routers.Execute(e)
}
