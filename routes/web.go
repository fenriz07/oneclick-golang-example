package routes

import (
	"github.com/fenriz07/oneclick-golang-example/controllers"
	"github.com/kataras/iris/v12"
)

/*SetWebRoutes defined web routes*/
func SetWebRoutes(app *iris.Application) {

	app.Get("/", controllers.Home)

	inscriptionRoutes := app.Party("/inscription")
	inscriptionRoutes.Get("/init", controllers.InitInscription)
	inscriptionRoutes.Post("/create", controllers.CreateInscription)
	inscriptionRoutes.Post("/confirm", controllers.ConfirmInscription)
	inscriptionRoutes.Get("/delete", controllers.DeleteInscriptionGet)
	inscriptionRoutes.Post("/delete", controllers.DeleteInscriptionPost)

	transactionRoutes := app.Party("/transaction")
	transactionRoutes.Get("/authorize", controllers.AuthorizeTransactionGet)
	transactionRoutes.Post("/authorize", controllers.AuthorizeTransactionPost)

	transactionRoutes.Get("/status", controllers.StatusTransactionGet)
	transactionRoutes.Post("/status", controllers.StatusTransactionPost)

	transactionRoutes.Get("/refund", controllers.RefundTransactionGet)
	transactionRoutes.Post("/refund", controllers.RefundTransactionPost)

}
