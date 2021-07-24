package controller

import (
	"aihealth/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// @Description View books detail information with Book ID
// @Accept  json
// @Produce  json
// @Param isbn path string true "Book ISBN"
// @Success 200
// @Header 200 string books "Books"
// @Router /books/isbn/{isbn} [get]
func GetBooksByISBN(c *gin.Context) {
	ISBN := c.Params.ByName("isbn")
	log.Println(ISBN)
	var books []model.Book

	result := model.MySQLPool.Where("isbn = ?", ISBN).Find(&books)
	// SELECT * FROM books WHERE id IN (1,2,3);

	if result.Error == nil {
		c.JSON(http.StatusOK, books)
	} else {
		log.Print(result.Error)
	}
}

// @Description View books detail information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 string books "Books"
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []model.Book

	result := model.MySQLPool.Find(&books)
	// SELECT * FROM books;

	if result.Error == nil {
		c.JSON(http.StatusOK, books)
	} else {
		log.Print(result.Error)
	}
}

// @Description Add books detail information
// @Accept json
// @Produce json
// @Param book body model.Book true "Add Book"
// @Success 200
// @Router /books [post]
func AddBooks(c *gin.Context) {
	log.Println(c)
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	RowsAffected, err := book.Insert()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	} else {
		c.JSON(200, gin.H{"RowsAffected": RowsAffected, "data": book})
	}
}

// @Description Delete book by isbn
// @Accept json
// @Produce json
// @Param  isbn path string true "ISBN"
// @Success 200
// @Router /books/{isbn} [delete]
func DeleteBookByISBN(c *gin.Context) {
	isbn := c.Params.ByName("isbn")
	log.Println(isbn)

	result := model.MySQLPool.Where("isbn = ?", isbn).Delete(&model.Book{})

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"RowsAffected": result.RowsAffected})
	} else {
		log.Println("error...", result)
		c.JSON(400, gin.H{"err": result.Error.Error()})
	}
}

// @Description Update book by isbn
// @Accept json
// @Produce json
// @Param  isbn path string true "Book Id"
// @Param book body model.Book true "Add Book"
// @Success 200
// @Router /books/{isbn} [patch]
func UpdateBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	mutex := model.RediSync.NewMutex("update books")

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		log.Fatal(err)
		panic(err)
	}
	// Do your work that requires the lock.
	result := model.MySQLPool.Model(&model.Book{}).Where("isbn = ?", isbn).Updates(book)
	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		log.Fatal(err)
		panic("unlock failed")
	}

	if result.Error != nil {
		// log.Fatal(err)
		c.JSON(400, gin.H{"Error Message": result.Error.Error()})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"RowsAffected": result.RowsAffected, "data": book})
	}
}
