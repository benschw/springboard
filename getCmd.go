package main

import (
	"context"
	"flag"
	"fmt"

	b64 "encoding/base64"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

type getCmd struct {
	vault   *vaultapi.Logical
	secrets string
	path    string
	key     string
	value   string
}

func (*getCmd) Name() string     { return "get" }
func (*getCmd) Synopsis() string { return "Encrypt a value and set it to the specified key" }
func (*getCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *getCmd) SetFlags(f *flag.FlagSet) {
	//	f.StringVar(&p.secrets, "secrets", "", "path to secrets file")
	f.StringVar(&p.path, "path", "", "vault path")
	f.StringVar(&p.value, "value", "", "plain text secret")

}

func (p *getCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	v, err := p.vault.Write("transit/decrypt/cub",
		map[string]interface{}{
			"ciphertext": p.value,
		})
	if err != nil {
		panic(err)
	}

	enc := v.Data["plaintext"]
	dec, err := b64.StdEncoding.DecodeString(fmt.Sprintf("%s", enc))
	fmt.Printf("%s\n", dec)

	return subcommands.ExitSuccess
}
