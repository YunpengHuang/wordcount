package services

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

type SqliteDB struct {
	DB *sql.DB
}

type WordValue struct {
	Text            string
	ValueDifference string
	Correct         bool
	Colour          bool
}

func CalcString(str *string) (int, error) {
	sum := 0
	for _, letter := range strings.ToUpper(*str) {
		value := int(letter) - 64
		sum += value
	}
	return sum, nil
}

func (lite *SqliteDB) GetWord(word string) *WordValue {
	wordValue, err := CalcString(&word)
	if err != nil {
		return nil
	}
	query := "SELECT word_of_the_day, word_value FROM wordcount WHERE id = (SELECT MAX(id) FROM wordcount) LIMIT 1"
	stmt, err := lite.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer stmt.Close()

	var dbValue int
	var dbWord string
	var isCorrect bool

	err = stmt.QueryRow().Scan(&dbWord, &dbValue)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	difference := dbValue - wordValue
	isCorrect = false

	// Colour false (red), true (green)
	colour := false
	if difference == 0 && dbWord == word {
		isCorrect = true
	} else if difference < 0 {
		colour = true
	}
	value := strconv.Itoa(difference)
	return &WordValue{
		Text:            word,
		ValueDifference: value,
		Correct:         isCorrect,
		Colour:          colour,
	}
}
