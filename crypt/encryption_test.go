package crypt

import (
	"testing"
	vaultapi "github.com/hashicorp/vault/api"
)

func TestEncryption(t *testing.T) {

	
	cfg := vaultapi.DefaultConfig()
	client, err := vaultapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	c := client.Logical()


	lib := New(c, "my-key")

	secret := "hello world"

	enc, err := lib.Encrypt(secret)
	if err != nil {
		t.Error(err)
	}

	dec, err := lib.Decrypt(enc)
	if err != nil {
		t.Error(err)
	}

	if dec != secret {
		t.Error("decrypted doesn't match original", dec)
	}
}
