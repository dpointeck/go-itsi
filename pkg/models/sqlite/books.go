package sqlite

import (
	"database/sql"

	"github.com/dpointeck/go-itsi/pkg/models"
)

// BookModel - Sqlite Model
type BookModel struct {
	DB *sql.DB
}

// Insert - This will insert a new snippet into the database.
func (m *BookModel) Insert(title, isbn, released string) (int, error) {

	stmt, err := m.DB.Prepare(`INSERT INTO books (title, isbn, released, created, updated) VALUES(?,?,DATETIME(?), CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(title, isbn, released)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get - This will return a specific Book based on its id.
func (m *BookModel) Get(id int) (*models.Book, error) {
	return nil, nil
}

// Latest - This will return the 10 most recently created snippets.
func (m *BookModel) Latest() ([]*models.Book, error) {
	return nil, nil
}
