package model

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// bookinfo = {}

const (
	_selectUser = "select * from user where id = ?"
	_updateUser = "update user set name = ? where name = ?"
	_insertUser = "insert user (name,age) values (?,?)"
	_deleteUser = "delete user where id = ?"
)

var db *sql.DB

type BookType struct {
	ISBN      string `json:"isbn"`
	Author    string `json:"author"`
	Title     string `json:"title" example:"Go in Action"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Tags      string `json:"tags"`
	Language  string `json:"language"`
	Rating    int8   `json:"rating"`
}

// Create
func (b *BookType) add_to(book BookType) error {
	// ...
	return nil
}

// Delete
func (b *BookType) delete_from(book BookType) error {
	// ...
	return nil
}

// Read
func readhandler(w http.ResponseWriter, r *http.Request) (*sql.Rows, error) {
	resultRows, err := db.Query("SELECT * FROM books")
	return resultRows, err
}

// Update
func (b *BookType) update() error {
	//db.Exec()
	return nil
}
