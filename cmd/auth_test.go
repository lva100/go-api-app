package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"go/adv-demo/pkg/hsrv"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginAccess(t *testing.T) {
	ts := httptest.NewServer(hsrv.App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "vasya@mail.ru",
		Password: "123",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token is empty")
	}
}

func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(hsrv.App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "vasya@mail.ru",
		Password: "1234",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}
}
