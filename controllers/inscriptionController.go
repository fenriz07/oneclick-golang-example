package controllers

import (
	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/inscription"
	"github.com/kataras/iris/v12"
)

/*CreateInscription controller for create inscription*/
func CreateInscription(ctx iris.Context) {

	response, err := inscription.CreateInscription("1212", "servio.za@gmail.com", "http://localhost:8080/inscription/confirm")

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
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
	} else {
		ctx.ViewData("authorizationCode", response.AuthorizationCode)
		ctx.ViewData("cardNumber", response.CardNumber)
		ctx.ViewData("cardType", response.CardType)
		ctx.ViewData("responseCode", response.ResponseCode)
		ctx.ViewData("tbkUser", response.TbkUser)

		ctx.View("confirm_inscription.html")
	}

}
