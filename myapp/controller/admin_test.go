package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdmLogin(t *testing.T) {
	url := "http://localhost:8000/login"

	var jsonStr = []byte(`{"email":"resort@gmail.com", "password":"qwertypoiu"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"message":"login successful"}`

	assert.JSONEq(t, expResp, string(body))
}

func TestAdmUserNotExist(t *testing.T) {
	url := "http://localhost:8000/login"

	var data = []byte(`{"email":"jp@gmail.com", "password":"pass"}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.JSONEq(t, `{"error": "sql: no rows in result set"}`, string(body))
}
