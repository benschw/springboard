package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Usage: springboard <subcommand> -s <secrets file> -t <transit key> [args]\n\n")

	fmt.Printf("Subcommands:\n")
	fmt.Printf("    help                display this help screen and exit\n")
	fmt.Printf("    set <key> <value>   set/encrypt 'value' in local secrets file\n")
	fmt.Printf("    get <key>           get/decrypt 'value' from local secrets file\n")
	fmt.Printf("    remove <key>        remove 'key' (and its value) from local secrets file\n")
	fmt.Printf("    push <path>         publish secrets in local secrets file to\n")
	fmt.Printf("                        'path' in vault generic secrets backend\n\n")

	fmt.Printf("Flags:\n")
	fmt.Printf("    -s string   secrets file\n")
	fmt.Printf("    -t string   transit key\n\n")

	fmt.Printf("Examples:\n")
	fmt.Printf("    springboard set -s secrets.yml -t my-key user_name supersecret\n")
	fmt.Printf("    springboard get -s secrets.yml -t my-key user_name\n")
	fmt.Printf("    springboard remove -s secrets.yml -t my-key user_name\n")
	fmt.Printf("    springboard push -s secrets.yml -t my-key secret/my-space\n\n")

	fmt.Printf("github.com/benschw/springboard\n")
}

func main() {
	// flags
	f := flag.NewFlagSet("", flag.ExitOnError)
	f.Usage = usage

	secretsFile := f.String("s", "", "secrets file path")
	transitKey := f.String("t", "", "transit key")

	if len(os.Args) < 2 {
		f.Usage()
		os.Exit(2)
	}

	f.Parse(os.Args[2:])

	if *secretsFile == "" || *transitKey == "" {
		f.Usage()
		os.Exit(2)
	}

	args := f.Args()

	// App
	app, err := NewApp(*secretsFile, *transitKey)
	if err != nil {
		f.Usage()
		os.Exit(1)
	}

	// subcommands
	switch os.Args[1] {
	case "get":
		if len(args) != 1 {
			f.Usage()
			os.Exit(2)
		}
		if err := app.get(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "set":
		if len(args) != 2 {
			f.Usage()
			os.Exit(2)
		}
		if err := app.set(args[0], args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "remove":
		if len(args) != 1 {
			f.Usage()
			os.Exit(2)
		}
		if err := app.remove(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "push":
		if len(args) != 1 {
			f.Usage()
			os.Exit(2)
		}
		if err := app.push(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "help":
		f.Usage()
		os.Exit(0)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		f.Usage()
		os.Exit(2)
	}

	os.Exit(0)
}
