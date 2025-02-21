package auth

import "testing"

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
