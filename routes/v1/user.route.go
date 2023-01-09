package v1

import (
	"iris/controllers"
	"iris/middleware"

	"github.com/kataras/iris/v12"
)

func UserRoutes(app *iris.Application) {
	v1 := app.Party("/v1", middleware.VerifyToken)
	v1.Get("/users", controllers.GetUsers)

}
