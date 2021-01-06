package controllers

import "github.com/kataras/iris/v12"

func Home(ctx iris.Context) {
	ctx.View("home.html")
}
