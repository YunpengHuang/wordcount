package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/YunpengHuang/wordcount/app/components"
	"github.com/YunpengHuang/wordcount/app/views/layout"

	"github.com/YunpengHuang/wordcount/services"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type WordList = []services.WordValue

func main() {
	db, err := sql.Open("sqlite3", "./wordcount.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := `CREATE TABLE IF NOT EXISTS wordcount (
            id INTEGER PRIMARY KEY,
            word_of_the_day TEXT NOT NULL,
            word_value INTEGER NOT NULL
  );`
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("%q %s\n", err, sql)
	}

	var wordList []services.WordValue
	var hasReachedLimit bool
	var hasCorrectGuess bool
	service := &services.SqliteDB{
		DB: db,
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		component := layout.Base(wordList, hasReachedLimit, hasCorrectGuess)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/result", func(c echo.Context) error {
		text := c.QueryParam("query")
		if text == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "bad bad")
		}
		valueResult := service.GetWord(text)
		wordList = append(wordList, *valueResult)
		if len(wordList) < 10 {
			hasReachedLimit = false
		} else {
			hasReachedLimit = true
		}
		if valueResult.Correct {
			hasCorrectGuess = true
		}
		component := components.Input(wordList, hasReachedLimit, hasCorrectGuess)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/reset", func(c echo.Context) error {
		wordList = []services.WordValue{}
		hasReachedLimit = false
		hasCorrectGuess = false
		c.Response().Header().Set("HX-Redirect", "/")
		component := layout.Base(wordList, hasReachedLimit, hasCorrectGuess)
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":3000"))
}
