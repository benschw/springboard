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

type setCmd struct {
	vault   *vaultapi.Logical
	secrets string
	transitKey    string
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
	if len(a) != 2 {
		fmt.Println(p.Usage())
		return subcommands.ExitUsageError
	}
	key, valSlice := a[0], a[1:]
	if p.secrets == "" {
		fmt.Println("-s must be set")
		return subcommands.ExitUsageError
	}
	if p.transitKey == "" {
		fmt.Println("-t must be set")
		return subcommands.ExitUsageError
	}

	value := ""
	for _, arg := range valSlice {
		value = fmt.Sprintf("%s%s", value, arg)
	}

	c := crypt.New(p.vault, p.transitKey)

	s, err := secrets.New(p.secrets, c)
	if err != nil {
		panic(err)
	}

	s.Set(key, value)
	if err = s.Save(); err != nil {
		panic(err)
	}

	return subcommands.ExitSuccess
}
