package controller

import (
	"database/sql"
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "Invalid json body")
	}
	defer r.Body.Close()

	SaveErr := course.Create()
	if SaveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, SaveErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "course added successfully"})
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	courseid, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
	}
	c := model.Course{Cid: courseid}
	getErr := c.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Student not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, c)
}

func getCourseId(courseIdParam string) (int64, error) {
	courseId, courseErr := strconv.ParseInt(courseIdParam, 10, 64)
	if courseErr != nil {
		return 0, courseErr
	}
	return courseId, nil
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	old_cid := mux.Vars(r)["cid"]
	old_courseid, idErr := getCourseId(old_cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var course model.Course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&course); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	updateErr := course.Update(old_courseid)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, "Course not found")
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJSON(w, http.StatusOK, course)
	}
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	cid := mux.Vars(r)["cid"]
	courseid, idErr := getCourseId(cid)
	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	c := model.Course{Cid: courseid}
	if err := c.Delete(); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "Deleted"})
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, getErr := model.GetAllCourses()
	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, getErr.Error())
		return
	}
	httpResp.RespondWithJSON(w, http.StatusOK, courses)
}
