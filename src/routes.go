package src

import (
	"src/src/http/controllers/books"

	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	app.Use(iris.Compression)
	books.BooksRoutes(app)
}
