package main

import (
	"liftlab/src/config"
	"liftlab/src/helpers"
	"liftlab/src/routers"
)

func main() {
	config.LoadConfig()
	helpers.ConnectToDB()
	app := routers.SetUpRouter()
	port := config.PORT
	app.Listen(":" + port)
}
