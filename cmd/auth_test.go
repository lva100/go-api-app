package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/hsrv"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1"), bcrypt.DefaultCost)
	db.Create(&user.User{
		Email:    "test@test.ru",
		Password: string(hashedPassword),
		Name:     "Test user",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "test@test.ru").
		Delete(&user.User{})
}

func TestLoginAccess(t *testing.T) {
	//Prepare data
	db := initDB()
	initData(db)
	ts := httptest.NewServer(hsrv.App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.ru",
		Password: "1",
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
	removeData(db)
}

func TestLoginFail(t *testing.T) {
	//Prepare data
	db := initDB()
	initData(db)

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
	removeData(db)
}
