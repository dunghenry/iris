package controllers

import "github.com/kataras/iris/v12"

func GetUsers(ctx iris.Context) {
	token := ctx.Request().Header.Get("token")

	ctx.JSON(iris.Map{
		"status": 200,
		"token":  token,
	})
}
