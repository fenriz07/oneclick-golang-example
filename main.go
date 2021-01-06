package main

import (
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/oneclickmall"
	"github.com/fenriz07/oneclick-golang-example/routes"
	"github.com/kataras/iris/v12"
)

func main() {

	oneclickmall.SetEnvironmentIntegration()

	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	routes.SetWebRoutes(app)

	app.Listen(":8080")
}
