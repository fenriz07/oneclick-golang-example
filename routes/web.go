package routes

import (
	"github.com/fenriz07/oneclick-golang-example/controllers"
	"github.com/kataras/iris/v12"
)

/*SetWebRoutes defined web routes*/
func SetWebRoutes(app *iris.Application) {

	app.Get("/", controllers.Home)

	inscriptionRoutes := app.Party("/inscription")
	inscriptionRoutes.Get("/create", controllers.CreateInscription)
	inscriptionRoutes.Post("/confirm", controllers.ConfirmInscription)

}
