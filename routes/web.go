package routes

import "github.com/kataras/iris/v12"

/*SetWebRoutes defined web routes*/
func SetWebRoutes() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

}
