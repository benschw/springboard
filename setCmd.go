package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)
type setCmd struct {
	secrets string
	vault string
	key string
	value string
}

func (*setCmd) Name() string     { return "set" }
func (*setCmd) Synopsis() string { return "Encrypt a value and set it to the specified key" }
func (*setCmd) Usage() string {
	return `print [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *setCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.secrets, "secrets", "", "path to secrets file")
	f.StringVar(&p.vault, "vault", "", "vault host address")
	f.StringVar(&p.key, "key", "", "vault key")
	f.StringVar(&p.value, "value", "", "plain text secret")

}

func (p *setCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	fmt.Printf("%s | %s | %s | %s\n", p.secrets, p.vault, p.key, p.value)

	return subcommands.ExitSuccess
}

