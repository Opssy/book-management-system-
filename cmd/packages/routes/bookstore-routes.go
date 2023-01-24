package routes

import (
	"book-management-system-/cmd/packages/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.Getbook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

func main() {

}
