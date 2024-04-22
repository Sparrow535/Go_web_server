package controller

import (
	"encoding/json"
	"myapp/model"
	"net/http"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&stud); err != nil {
		w.Write([]byte("Invalid JSON data"))
		return
	}
	defer r.Body.Close()

	saveErr := stud.Create()
	if saveErr != nil {
		w.Write([]byte("Database error"))
		return
	}

	w.Write([]byte("response success."))
}
