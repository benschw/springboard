package main

import (
	"context"
	"flag"
	"fmt"

	b64 "encoding/base64"

	"github.com/google/subcommands"
	vaultapi "github.com/hashicorp/vault/api"
)

type setCmd struct {
	vault   *vaultapi.Logical
	secrets string
	path    string
	key     string
	value   string
}

func (*setCmd) Name() string     { return "set" }
func (*setCmd) Synopsis() string { return "Encrypt a value and set it to the specified key" }
func (*setCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *setCmd) SetFlags(f *flag.FlagSet) {
	//	f.StringVar(&p.secrets, "secrets", "", "path to secrets file")
	f.StringVar(&p.path, "path", "", "vault path")
	f.StringVar(&p.key, "key", "", "secret key")
	f.StringVar(&p.value, "value", "", "plain text secret")

}

func (p *setCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	enc := b64.StdEncoding.EncodeToString([]byte(p.value))

	v, err := p.vault.Write("transit/encrypt/cub",
		map[string]interface{}{
			"plaintext": enc,
		})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", v.Data["ciphertext"])

	return subcommands.ExitSuccess
}
