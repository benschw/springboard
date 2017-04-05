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
func (*setCmd) Synopsis() string { return "Encrypt a value and set it to the specified key" }
func (*setCmd) Usage() string {
	return "set -secrets <path to secrets.yml> -transit-key <transit key name> <key/name of secret> <secret to store>:"
}

func (p *setCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "secrets", "", "path to secrets file")
	f.StringVar(&p.transitKey, "transit-key", "", "vault transit key")
}

func (p *setCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	a := f.Args()
	key, valSlice := a[0], a[1:]

	value := ""
	for _, arg := range valSlice {
		value = fmt.Sprintf("%s%s", value, arg)
	}


	s, err := secrets.New(p.secrets)
	if err != nil {
		panic(err)
	}

	c := crypt.New(p.vault, p.transitKey)

	val, err := c.Encrypt(value)
	if err != nil {
		panic(err)
	}

	s.Set(key, val)
	if err = s.Save(); err != nil {
		panic(err)
	}

	return subcommands.ExitSuccess
}
