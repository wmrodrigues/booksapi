package book

import (
	"common"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"repositories"
	"structs"
	"time"
)

// Save creates a new book on database or updates one on database
// if it already exists
func Save(book structs.Book) (int64, error) {
	var err error
	var result int64

	bookRepository := repositories.NewBookRepository()
	defer bookRepository.Close()

	if book.Title == "" {
		err = errors.New("title attribute is required")
	}

	if book.Isbn == "" {
		err = errors.New("isbn attribute is required")
	}

	if book.ReleaseDate == nil {
		err = errors.New("release_date attribute is required")
	}

	_, err = time.Parse(common.DateFormat, fmt.Sprint(book.ReleaseDate))

	if err != nil {
		err = errors.New("release_date attribute must be on format yyyy-mm-dd")
	}

	if book.AuthorId == 0 {
		err = errors.New("author_id attribute is required and must be greater than zero")
	}

	if err != nil {
		return 0, err
	}

	if book.Id < 1 {
		result, err = bookRepository.Create(book)
	} else {
		result = book.Id
		err = bookRepository.Update(book)
	}

	if err != nil {
		log.Println(err)
	}

	return result, err
}

// Remove removes a specific book from database
func Remove(id int64) error {
	var err error
	bookRepository := repositories.NewBookRepository()
	defer bookRepository.Close()

	if id < 1 {
		err = errors.New("id param is required and must be greater than zero")
		return err
	}

	err = bookRepository.Delete(id)

	if err != nil {
		log.Println(err)
	}

	return err
}

// GetById get a specific book on database
func GetById(id int64) (structs.Book, error) {
	var err error
	var book structs.Book

	bookRepository := repositories.NewBookRepository()
	defer bookRepository.Close()

	if id < 1 {
		err = errors.New("id param is required and must be greater than zero")
		return book, err
	}

	result, err := bookRepository.GetById(id)

	if err != nil {
		log.Println(err)
	}

	bytes, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
		return book, err
	}

	err = json.Unmarshal(bytes, &book)

	if err != nil {
		log.Println(err)
		return book, err
	}

	return book, nil
}

// GetPaged gets a set of books on database
func GetPaged(limit, offset int) ([]structs.Book, error) {
	var err error
	var books []structs.Book

	bookRepository := repositories.NewBookRepository()
	defer bookRepository.Close()

	if limit < 1 {
		err = errors.New("limit param is required and must be greater than zero")
		return books, err
	}

	if offset < 0 {
		err = errors.New("offset param is required and must be greater than -1")
	}

	offset *= limit

	result, err := bookRepository.GetPaged(limit, offset)

	if err != nil {
		log.Println(err)
		return books, err
	}

	bytes, err := json.Marshal(result)

	if err != nil {
		log.Println(err)
		return books, err
	}

	err = json.Unmarshal(bytes, &books)

	if err != nil {
		log.Println(err)
		return books, err
	}

	return books, nil
}
