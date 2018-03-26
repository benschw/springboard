package main

import (
	"fmt"

	"github.com/benschw/springboard/crypt"
	"github.com/benschw/springboard/publisher"
	"github.com/benschw/springboard/secrets"
	vaultapi "github.com/hashicorp/vault/api"
)

func NewApp(secretsFile string, transitKey string) (*App, error) {
	cfg := vaultapi.DefaultConfig()
	client, err := vaultapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	vault := client.Logical()

	c := crypt.New(vault, transitKey)

	s, err := secrets.New(secretsFile, c)
	if err != nil {
		return nil, err
	}
	return &App{vault: vault, secrets: s}, nil
}

type App struct {
	vault   *vaultapi.Logical
	secrets *secrets.Secrets
}

func (a *App) get(key string) error {
	val, err := a.secrets.Get(key)
	if err != nil {
		return err
	}

	fmt.Println(val)
	return nil
}

func (a *App) set(key string, value string) error {
	if err := a.secrets.Set(key, value); err != nil {
		return err
	}
	return a.secrets.Save()
}

func (a *App) remove(key string) error {
	if err := a.secrets.Remove(key); err != nil {
		return err
	}
	return a.secrets.Save()
}

func (a *App) push(path string) error {
	pub := publisher.New(a.vault, path)

	return pub.Push(a.secrets)
}
