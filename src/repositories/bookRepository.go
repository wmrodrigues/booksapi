package repositories

import (
	"config"
	"db"
	"errors"
	"log"
	"structs"
)

// BookRepository represents book table
type BookRepository struct {
	db db.Database
}

// NewBookRepository creates a new instance of BookRepository
func NewBookRepository() *BookRepository {
	config := db.Config{Host: config.CONFIG.Books.Host,
		Port:     config.CONFIG.Books.Port,
		User:     config.CONFIG.Books.User,
		Password: config.CONFIG.Books.Password,
		Database: config.CONFIG.Books.Database}
	conn := db.NewPgSQL(config)
	return &BookRepository{db: conn}
}

// Create inserts a book record on book database table
func (b *BookRepository) Create(book structs.Book) (int64, error) {
	var id int64
	query := b.createQuery()
	err := b.db.QueryRow(query, book.Title,
		book.Isbn,
		book.About,
		book.Edition,
		book.PageNumber,
		book.ReleaseDate,
		book.AuthorId).Scan(&id)

	if err != nil {
		err = errors.New("error inserting book on database, " + err.Error())
		log.Println(err)
		return 0, err
	}

	return id, nil
}

// Update updates a book record on book database table
func (b *BookRepository) Update(book structs.Book) error {
	query := b.updateQuery()
	_, err := b.db.Execute(query, book.Title,
		book.Isbn,
		book.About,
		book.Edition,
		book.PageNumber,
		book.ReleaseDate,
		book.AuthorId,
		book.Id)

	if err != nil {
		err = errors.New("error updating book on database, " + err.Error())
		log.Println(err)
		return err
	}

	return nil
}

// Delete delete a book from database based on id attribute
func (b *BookRepository) Delete(id int64) error {
	query := b.deleteQuery()
	_, err := b.db.Execute(query, id)

	if err != nil {
		err = errors.New("error deleting book from database, " + err.Error())
		log.Println(err)
		return err
	}

	return nil
}

// GetById get a book record on database based on id attribute
func (b *BookRepository) GetById(id int64) (map[string]interface{}, error) {
	query := b.getByIdQuery(id)
	result, err := b.db.MapScan(query)

	if err != nil {
		err = errors.New("error getting book from database, " + err.Error())
		log.Println(err)
		return nil, err
	}

	return result, nil
}

// GetPaged get a set of books on database limited on limit param starting on
// offset param
func (b *BookRepository) GetPaged(limit, offset int) ([]map[string]interface{}, error) {
	query := b.getPagedQuery(limit, offset)
	result, err := b.db.SliceMap(query)

	if err != nil {
		err = errors.New("error getting book set from database, " + err.Error())
		log.Println(err)
		return nil, err
	}

	return result, nil
}

// Close closes the database connection
func (b *BookRepository) Close() {
	b.db.Close()
}
