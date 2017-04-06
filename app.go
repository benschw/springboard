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

func (a *App) get(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("missing argument")
	}
	key := args[0]

	val, err := a.secrets.Get(key)
	if err != nil {
		return err
	}

	fmt.Println(val)
	return nil
}

func (a *App) set(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("missing argumeant")
	}
	key, value := args[0], args[1]

	if err := a.secrets.Set(key, value); err != nil {
		return err
	}
	return a.secrets.Save()
}

func (a *App) push(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("missing argument")
	}
	path := args[0]

	pub := publisher.New(a.vault, path)

	return pub.Push(a.secrets)
}
