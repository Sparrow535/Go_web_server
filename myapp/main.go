package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/home", HomeHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		fmt.Println("error: ", err)
	}
}
