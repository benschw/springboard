package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/benschw/vault-cub/secrets"
	"github.com/benschw/vault-cub/crypt"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

type getCmd struct {
	vault   *vaultapi.Logical
	secrets string
	transitKey string
}

func (*getCmd) Name() string     { return "get" }
func (*getCmd) Synopsis() string { return "Decrypt a value from a secrets file" }
func (*getCmd) Usage() string    { return "get -s <./secrets.yml> -t <my-key> <key/name of secret>:\n" }

func (p *getCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "s", "", "path to secrets file")
	f.StringVar(&p.transitKey, "t", "", "vault transit key")
}

func (p *getCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	a := f.Args()
	if len(a) < 1 || p.secrets == "" || p.transitKey == "" {
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	key := a[0]

	c := crypt.New(p.vault, p.transitKey)

	s, err := secrets.New(p.secrets, c)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	
	val, err := s.Get(key)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println(val)

	return subcommands.ExitSuccess
}
