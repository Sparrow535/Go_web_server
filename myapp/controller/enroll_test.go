package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnroll(t *testing.T) {
	url := "http://localhost:8000/enroll"

	var jsonStr = []byte(`{"stdid": 1001, "cid": "103"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "my-cookie", Value: "my-value"})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	expResp := `{"status": "enrolled"}`
	assert.JSONEq(t, expResp, string(body))
}

func TestGetEnroll(t *testing.T) {
	c := http.Client{}
	url := "http://localhost:8000/enroll/1001/103"
	resp, _ := c.Get(url)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	expResp := `{"stdid": 1001, "cid": "103", "date": "2024-05-07T07:09:32Z"}`
	assert.JSONEq(t, expResp, string(body))
}
