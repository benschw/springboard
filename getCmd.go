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
func (*getCmd) Synopsis() string { return "Encrypt a value and set it to the specified key" }
func (*getCmd) Usage() string {
	return "get -secrets <path to secrets.yml> -transit-key <transit key name> <key/name of secret>:"

}

func (p *getCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "secrets", "", "path to secrets file")
	f.StringVar(&p.transitKey, "transit-key", "", "vault transit key")
}

func (p *getCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	a := f.Args()
	key := a[0]

	s, err := secrets.New(p.secrets)
	if err != nil {
		panic(err)
	}
	
	str, err := s.Get(key)
	if err != nil {
		panic(err)
	}

	c := crypt.New(p.vault, p.transitKey)

	val, err := c.Decrypt(str)
	if err != nil {
		panic(err)
	}

	fmt.Println(val)

	return subcommands.ExitSuccess
}
