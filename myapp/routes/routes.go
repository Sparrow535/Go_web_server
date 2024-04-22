package routes

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	var port = 8080
	router := mux.NewRouter()
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	log.Println("Application running on port", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
