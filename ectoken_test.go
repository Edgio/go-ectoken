package ectoken

import (
	"testing"
)

const (
	key   = "somekey"
	token = "ec_expire=1257642471&ec_secure=33"
	ermsg = "expected string hash recieved nothing."
)

func TestEncrypt(t *testing.T) {
	// should always get a hash back greater than 0
	if result := len(Encrypt(key, "")); result <= 0 {
		t.Error(ermsg)
	}

	if result := len(Encrypt(key, token)); result <= 0 {
		t.Error(ermsg)
	}
}

func TestDecrypt(t *testing.T) {
	tokenHash := Encrypt(key, token)

	result, err := Decrypt(key, tokenHash)
	if err != nil {
		t.Error(err)
	}

	// the decrypted token value should equal the string used to encrypt
	if result != token {
		t.Errorf("expected decrypted value to be %s, got %s\n", token, tokenHash)
	}
}
