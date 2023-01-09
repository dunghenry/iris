package middleware

import "github.com/kataras/iris/v12"

func VerifyToken(ctx iris.Context) {
	ctx.Request().Header.Add("token", "Bearer token")
	ctx.Next()
}
