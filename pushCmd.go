package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/benschw/vault-cub/publisher"
	"github.com/benschw/vault-cub/crypt"
	"github.com/benschw/vault-cub/secrets"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

type pushCmd struct {
	vault   *vaultapi.Logical
	secrets string
	transitKey string
}

func (*pushCmd) Name() string     { return "push" }
func (*pushCmd) Synopsis() string { return "Dencrypt a value from a secrets file" }
func (*pushCmd) Usage() string {
	return "get -s <./secrets.yml> -p <secret/my-space>:\n"
}

func (p *pushCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "s", "", "path to secrets file")
	f.StringVar(&p.transitKey, "t", "", "vault transit key")
}

func (p *pushCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	a := f.Args()
	if len(a) != 1 {
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	path := a[0]
	if p.secrets == "" {
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	if p.transitKey == "" {
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}

	c := crypt.New(p.vault, p.transitKey)

	s, err := secrets.New(p.secrets, c)
	if err != nil {
		panic(err)
	}

	pub := publisher.New(p.vault, path)

	if err = pub.Push(s); err != nil {
		panic(err)
	}


	return subcommands.ExitSuccess
}
