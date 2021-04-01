package controller

import (
	"fmt"
	"net/http"
)

// An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if r.URL.Path[1:] == "read" {
		fmt.Fprintf(w, "The books infomation: %s", r.URL.Path[1:])
	}
}

// func main() {
// 	var err error
// 	db, err = sql.Open("mysql", "root:966841@tcp(127.0.0.1:3306)/bookstore")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer db.Close()
// 	// which tells the http package to handle all requests to the web root ("/") with handler.
// 	// http.HandleFunc("/", handler)
// 	http.HandleFunc("/read", readhandler)
// 	log.Fatal(http.ListenAndServe(":8090", nil))
// }
