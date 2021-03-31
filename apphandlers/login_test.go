package apphandlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	username := "c137@onecause.com"
	password := "#th@nH@rm#y#r!$100%D0p#"
	ott := "0305"
	dataToSend := fmt.Sprintf(`{"username": "%s", "password": "%s", "ott": "%s"}`, username, password, ott)
	payload := []byte(dataToSend)
	req, _ := http.NewRequest(http.MethodPost, "/apiv1/login", bytes.NewBuffer(payload))
	handler := http.HandlerFunc(LoginHandler)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	status := res.Code

	if status != http.StatusOK {
		t.Errorf("Wrong response code. Got: %v  Expecting: %v", status, http.StatusOK)
	}

}
