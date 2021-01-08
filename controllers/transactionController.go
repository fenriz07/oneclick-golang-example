package controllers

import (
	"strconv"

	"github.com/fenriz07/Golang-Transbank-Oneclick-mall/pkg/transaction"
	"github.com/kataras/iris/v12"
)

func AuthorizeTransactionGet(ctx iris.Context) {
	ctx.View("authorize_transaction_get.html")
}

func AuthorizeTransactionPost(ctx iris.Context) {

	username := ctx.PostValue("username")
	userToken := ctx.PostValue("user_token")
	amountStr := ctx.PostValue("amount")
	order := ctx.PostValue("order")

	amount, err := strconv.Atoi(amountStr)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	}

	detail := transaction.CreateDetailTransaction("597055555542", order, amount, 1)

	response, err := transaction.AuthorizeTransaction(username, userToken, order, detail)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	}

	ctx.ViewData("response", response)
	ctx.View("authorize_transaction_post.html")

}

func StatusTransactionGet(ctx iris.Context) {
	ctx.View("status_transaction_get.html")
}

func StatusTransactionPost(ctx iris.Context) {
	order := ctx.PostValue("order")

	response, err := transaction.StatusTransaction(order)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	}

	ctx.ViewData("response", response)
	ctx.View("status_transaction_post.html")

}

func RefundTransactionGet(ctx iris.Context) {
	ctx.View("refund_transaction_get.html")
}

func RefundTransactionPost(ctx iris.Context) {

	amountStr := ctx.PostValue("amount")
	order := ctx.PostValue("order")

	amount, err := strconv.Atoi(amountStr)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")
		return
	}

	response, err := transaction.RefundTransaction(order, "597055555542", order, amount)

	if err != nil {
		ctx.ViewData("error", err)
		ctx.View("error.html")

		return
	}

	ctx.ViewData("response", response)
	ctx.View("refund_transaction_post.html")
}
