package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/benschw/springboard/crypt"
	"github.com/benschw/springboard/secrets"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

type setCmd struct {
	vault      *vaultapi.Logical
	secrets    string
	transitKey string
}

func (*setCmd) Name() string     { return "set" }
func (*setCmd) Synopsis() string { return "Encrypt a value and store it to a secrets file" }
func (*setCmd) Usage() string {
	return "set -s <./secrets.yml> -t <my-key> <key/name of secret> <secret to store>:\n"
}

func (p *setCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "s", "", "path to secrets file")
	f.StringVar(&p.transitKey, "t", "", "vault transit key")
}

func (p *setCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	a := f.Args()
	if len(a) != 2 || p.secrets == "" || p.transitKey == "" {
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	key, value := a[0], a[1]

	c := crypt.New(p.vault, p.transitKey)

	s, err := secrets.New(p.secrets, c)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if err = s.Set(key, value); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	if err = s.Save(); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
