package crypt

import (
	"fmt"
	b64 "encoding/base64"
	vaultapi "github.com/hashicorp/vault/api"
)

func New(vault *vaultapi.Logical, key string) *Crypt  {
	return &Crypt{vault: vault, key: key}
}

type Crypt struct {
	vault *vaultapi.Logical
	key string
}

func (c *Crypt) Encrypt(val string) (string, error) {
	enc := b64.StdEncoding.EncodeToString([]byte(val))

	v, err := c.vault.Write(fmt.Sprintf("transit/encrypt/%s", c.key),
		map[string]interface{}{
			"plaintext": enc,
		})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", v.Data["ciphertext"]), nil
}

func (c *Crypt) Decrypt(val string) (string, error) {
	v, err := c.vault.Write(fmt.Sprintf("transit/decrypt/%s", c.key),
		map[string]interface{}{
			"ciphertext": val,
		})
	if err != nil {
		return "", err
	}

	enc := v.Data["plaintext"]
	dec, err := b64.StdEncoding.DecodeString(fmt.Sprintf("%s", enc))

	if err != nil {
		return "", err
	}
	return string(dec[:]), nil
}
