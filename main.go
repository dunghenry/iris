package main

import (
	v1 "iris/routes/v1"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

const maxSize = 8 * iris.MB

func main() {
	app := iris.New()
	config := iris.HTML("./views", ".html")
	app.RegisterView(config)
	app.HandleDir("/public", iris.Dir("./public"))
	v1.UserRoutes(app)
	// f, _ := os.Create("logger.log")
	// app.Logger().SetOutput(f)
	app.Get("/", func(ctx iris.Context) {
		// ctx.WriteString("Hi")
		// ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  200,
			"message": "Hello, world!",
		})
	})

	app.Get("/{id}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		data, _ := strconv.Atoi(id)
		ctx.JSON(iris.Map{
			"status": 200,
			"id":     data,
		})
	})
	app.Get("/home", func(ctx iris.Context) {
		ctx.ViewData("msg", "Xin chao")
		ctx.View("index")
	})
	app.Post("/upload", func(ctx iris.Context) {
		ctx.SetMaxRequestBodySize(maxSize)
		file, fileHeader, err := ctx.FormFile("image")
		// fmt.Println(file, fileHeader)
		if err != nil {
			ctx.JSON(iris.Map{
				"status":     400,
				"message":    err,
				"file":       file,
				"fileHeader": fileHeader,
			})
		} else {
			data := time.Now().Unix()
			filename := strconv.Itoa(int(data)) + "." + strings.Split(fileHeader.Filename, ".")[1]
			dest := filepath.Join("./public/images/" + filename)
			ctx.SaveFormFile(fileHeader, dest)
			ctx.JSON(iris.Map{
				"status":   200,
				"filename": filename,
			})
		}

	})
	app.Listen(":8080")
}
