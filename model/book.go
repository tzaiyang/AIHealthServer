package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ISBN      string `json:"isbn"`
	Author    string `json:"author"`
	Title     string `json:"title" example:"Go in Action"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	Tags      string `json:"tags"`
	Language  string `json:"language"`
	Rating    int8   `json:"rating"`
}

// Read
// func readhandler(w http.ResponseWriter, r *http.Request) (*sql.Rows, error) {

// }

// Query the database for the information requested and prints the results.
// If the query fails exit the program with an error.
// func (book Book) Query(ctx context.Context, id int64) (*sql.Rows, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancel()

// 	var name string
// 	err := MySQLPool.QueryRowContext(ctx, "select p.name from people as p where p.id = :id;", sql.Named("id", id)).Scan(&name)
// 	if err != nil {
// 		log.Fatal("unable to execute search query", err)
// 	}
// 	log.Println("name=", name)

// 	resultRows, err := MySQLPool.Query("SELECT * FROM books")
// 	return resultRows, err
// }

func (book Book) Insert() (int64, error) {
	result := MySQLPool.Create(&book)
	log.Println("Insert data: ", book)
	return result.RowsAffected, result.Error
}
