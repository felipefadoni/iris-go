package books

import (
	"github.com/kataras/iris/v12"
)

func BooksRoutes(app *iris.Application) {
	booksAPI := app.Party("/books")
	{
		// GET: http://localhost:8080/books
		booksAPI.Get("/", ListBooks)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", CreateBook)
	}
}
