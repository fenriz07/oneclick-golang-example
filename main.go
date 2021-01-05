package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	routes.SetWebRoutes()

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Listen(":8080")
}
