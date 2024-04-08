package controller

import (
	"encoding/json"
	"fmt"
	"myapp/model"
	"net/http"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	var stud model.Student

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&stud); err != nil {
		w.Write([]byte("invalid json data"))
		return
	}
	defer r.Body.Close()

	saveErr := stud.Create()
	if saveErr != nil {
		fmt.Println(saveErr)
		w.Write([]byte("Database ErrorSE"))
		return
	}
	w.Write([]byte("saved the student data"))

	// no error
}
