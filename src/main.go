package main

import "github.com/SomeBottle/MCSounds/router"

func main() {
	app := router.InitRouter()
	app.Run(":8080")
}
