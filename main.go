package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

func main() {


	cfg := vaultapi.DefaultConfig()
	client, err := vaultapi.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	c := client.Logical()


	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(&setCmd{vault: c}, "")
	subcommands.Register(&getCmd{vault: c}, "")
	subcommands.Register(&pushCmd{vault: c}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
