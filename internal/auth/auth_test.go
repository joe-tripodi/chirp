package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestPasswords(t *testing.T) {
	pwd := "password"
	hpwd, err := HashPassword(pwd)
	if err != nil {
		t.Error("unable to has simple password")
	}
	ok := CheckPasswordHash(pwd, hpwd)
	if ok != nil {
		t.Error("passwords do not match when they should")
	}

	ok = CheckPasswordHash("notpassword", hpwd)
	if ok == nil {
		t.Error("passwords should not match!!!!!")
	}
}

func TestJWTMake(t *testing.T) {
	const tokenSecret = "SlapTheBase"

	id, _ := uuid.NewRandom()
	ss, err := MakeJWT(id, tokenSecret, time.Duration(10*time.Minute))
	if err != nil {
		t.Errorf("Unable to generate token: %s\n", err)
	}

	userId, err := ValidateJWT(ss, tokenSecret)
	if userId != id {
		t.Errorf("expected ids to match: want: %v, got: %v\n", id, userId)
	}
}

func TestGetBearerToken(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"Bearer TOKEN_STRING"},
	}

	authToken, err := GetBearerToken(header)
	if err != nil {
		t.Errorf("Expected to get Auth Token but got err: %s", err)
	}
	if authToken == "" {
		t.Error("Auth Token is empty, this should not be the case")
	}
	if authToken != "TOKEN_STRING" {
		t.Errorf("Expected Auth Token to be: %s, but got: %s", "TOKEN_STRING", authToken)
	}
}

func TestGetBearerTokenBadPath(t *testing.T) {
	header := http.Header{
		"": []string{},
	}
	_, err := GetBearerToken(header)
	if err != ErrNoAuthHeader {
		t.Errorf("Expected ErrNoBearerToken")
	}
}

func TestGetBearerTokenBadHeader(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"Bearerr TOKEN_STRING"},
	}
	_, err := GetBearerToken(header)
	if err != ErrInvalidBearerToken {
		t.Errorf("Expected ErrInvalidBearerToken")
	}
}
