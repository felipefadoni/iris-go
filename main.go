package main

import (
	"src/src"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	src.Routes(app)
	app.Listen(":8080")
}
