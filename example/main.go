package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Book struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func createBook(db *sqlx.DB) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		_, err := db.Exec(
			"INSERT INTO books(id, name) VALUES(nextval('books_seq'), $1)",
			c.Param("name"),
		)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusOK)
	}
}

func listBooks(db *sqlx.DB) func(c *echo.Context) error {
	return func(c *echo.Context) error {
		var books []Book
		err := db.Select(&books, "SELECT * FROM books")
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.JSON(http.StatusOK, books)
	}
}

func NewConnectionPool(url string) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func Server(db *sqlx.DB) *echo.Echo {
	e := echo.New()
	e.Post("/books/:name", createBook(db))
	e.Get("/books", listBooks(db))
	return e
}

func main() {
	db := NewConnectionPool("postgres://localhost:5432/example?sslmode=disable")
	e := Server(db)
	e.Run(":3000")
}
