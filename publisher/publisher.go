package publisher

import (
	vaultapi "github.com/hashicorp/vault/api"
)

type secretsStore interface {
	Keys() []string
	Get(string) (string, error)
}

func New(vault *vaultapi.Logical, path string) *Publisher {
	return &Publisher{
		vault: vault,
		path: path,
	}
}

type Publisher struct {
	vault *vaultapi.Logical
	path string
}

func (p *Publisher) Push(secrets secretsStore) error {

	dec := make(map[string]interface{})

	for _, key := range secrets.Keys() {
		val, err := secrets.Get(key)
		if err != nil {
			return err
		}
		dec[key] = val
	}

	_, err := p.vault.Write(p.path, dec)

	return err
}
