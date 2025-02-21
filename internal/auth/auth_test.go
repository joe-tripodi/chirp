package auth

import (
	"fmt"
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
	fmt.Print(ss)

	userId, err := ValidateJWT(ss, tokenSecret)
	if userId != id {
		t.Errorf("expected ids to match: want: %v, got: %v\n", id, userId)
	}

}
