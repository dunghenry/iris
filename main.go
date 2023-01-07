package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		// ctx.WriteString("Hi")
		// ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  200,
			"message": "Hello, world!",
		})
	})
	app.Listen(":8080")
}
