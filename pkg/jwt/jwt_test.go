package jwt_test

import (
	"go/adv-demo/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "test@test.ru"
	jwtService := jwt.NewJWT("]}pj,xw;+R=/[{tjuC}{S|,sL:3lk}nYCUAa-Z)r H?8eAjAn?Z5.d4.X-fQ-fgC")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
