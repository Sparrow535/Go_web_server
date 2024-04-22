package model

import (
	"myapp/dataStore/postgres"

	_ "github.com/lib/pq"
)

type Student struct {
	StdId     int64  `json:"std_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

const queryInsertUser = "INSERT INTO student(std_id, first_name, last_name, email) VALUES(123, John, Cena, JC@gmail.com);"

func (s *Student) Create() error {
	_, err := postgres.Db.Exec(queryInsertUser, s.StdId, s.FirstName, s.LastName, s.Email)
	return err
}
