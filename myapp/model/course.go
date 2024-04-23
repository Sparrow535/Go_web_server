package model

import "myapp/dataStore/postgres"

type Course struct {
	Cid        int64  `json:"cid"`
	Coursename string `json:"coursename"`
}

const queryInsertcourse = "INSERT INTO course(cid, coursename) VALUES($1, $2)"

func (c *Course) Create() error {
	_, err := postgres.Db.Exec(queryInsertcourse, c.Cid, c.Coursename)
	return err
}

const queryGetCourse = "SELECT cid, coursename FROM course WHERE cid=$1;"

func (c *Course) Read() error {
	return postgres.Db.QueryRow(queryGetCourse, c.Cid).Scan(&c.Cid, &c.Coursename)
}

const queryUpdateCourse = "UPDATE course SET cid=$1, coursename=$2 WHERE cid=$3 RETURNING cid;"

func (c *Course) Update(oldID int64) error {
	err := postgres.Db.QueryRow(queryUpdateCourse, c.Cid, c.Coursename, oldID).Scan(&c.Cid)
	return err
}

const queryDeleteCourse = "DELETE FROM course WHERE cid=$1 RETURNING cid;"

func (c *Course) Delete() error {
	if err := postgres.Db.QueryRow(queryDeleteCourse, c.Cid).Scan(&c.Cid); err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]Course, error) {
	rows, getErr := postgres.Db.Query("SELECT * FROM course;")
	if getErr != nil {
		return nil, getErr
	}
	courses := []Course{}
	for rows.Next() {
		var c Course
		dbErr := rows.Scan(&c.Cid, &c.Coursename)
		if dbErr != nil {
			return nil, dbErr
		}
		courses = append(courses, c)
	}
	rows.Close()
	return courses, nil
}
