package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

func main() {

	vault := os.Getenv("VAULT_ADDR")
	token := os.Getenv("VAULT_TOKEN")

	cfg := vaultapi.DefaultConfig()
	cfg.Address = vault
	client, err := vaultapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	client.SetToken(token)
	c := client.Logical()

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&setCmd{vault: c}, "")
	subcommands.Register(&getCmd{vault: c}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
