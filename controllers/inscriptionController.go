package controllers

import (
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/inscription"
	"github.com/kataras/iris/v12"
)

func InitInscription(ctx iris.Context) {
	ctx.View("init_inscription.html")
}

/*CreateInscription controller for create inscription*/
func CreateInscription(ctx iris.Context) {

	username := ctx.PostValue("username")
	email := ctx.PostValue("email")

	response, err := inscription.CreateInscription(username, email, "http://localhost:8080/inscription/confirm")

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	} else {
		ctx.ViewData("token", response.Token)
		ctx.ViewData("urloneclick", response.URLWebpay)

		ctx.View("create_inscription.html")
	}

}

//ConfirmInscription controller for confirm inscription
func ConfirmInscription(ctx iris.Context) {
	token := ctx.PostValue("TBK_TOKEN")

	response, err := inscription.ConfirmInscription(token)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	} else {
		ctx.ViewData("authorizationCode", response.AuthorizationCode)
		ctx.ViewData("cardNumber", response.CardNumber)
		ctx.ViewData("cardType", response.CardType)
		ctx.ViewData("responseCode", response.ResponseCode)
		ctx.ViewData("tbkUser", response.TbkUser)

		ctx.View("confirm_inscription.html")
	}

}

func DeleteInscriptionGet(ctx iris.Context) {
	ctx.View("delete_inscription_get.html")
}

func DeleteInscriptionPost(ctx iris.Context) {

	username := ctx.PostValue("username")
	userToken := ctx.PostValue("user_token")

	status, err := inscription.DeleteInscription(userToken, username)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	}

	ctx.ViewData("status", status)
	ctx.View("delete_inscription_post.html")
}
