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
		response, _ := json.Marshal(map[string]string{"error": "invalid json body"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}
	defer r.Body.Close()

	saveErr := stud.Create()
	if saveErr != nil {
		response, _ := json.Marshal(map[string]string{"error": saveErr.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(map[string]string{"status": "student added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
