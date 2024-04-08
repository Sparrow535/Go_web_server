package route

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()
	// router.HandleFunc("/students/{course}", controller.Home)
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
