package main

import (
	"github.com/celt237/iris-enhance"
	"github.com/celt237/iris-enhance.demo/internal"
	"github.com/celt237/iris-enhance.demo/internal/handler"
	"github.com/celt237/iris-enhance.demo/internal/service"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to Iris!</h1>")
	})
	iris_enhance.RegisterSwaggerDoc(app, "docs/swagger.json", "/doc")
	println("swagger doc: http://localhost:8081/doc/index")
	demoService := service.NewDemoService()
	apiHandler := &internal.ApiHandlerImpl{}
	handler.RegisterDemoHTTPHandler(app.Party("/demo"), apiHandler, demoService)
	app.Run(iris.Addr(":8081"))
}
