package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Usage: springboard <subcommand> -s <secrets file> -t <transit key> [values]\n\n")

	fmt.Printf("Flags:\n")
	fmt.Printf("    -s string\n")
	fmt.Printf("        secrets file\n")
	fmt.Printf("    -t string\n")
	fmt.Printf("        transit key\n\n")

	fmt.Printf("Examples:\n")
	fmt.Printf("    springboard set -s secrets.yml -t my-key user_name supersecret\n")
	fmt.Printf("    springboard get -s secrets.yml -t my-key user_name\n")
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
	args := f.Args()

	if *secretsFile == "" {
		f.Usage()
		os.Exit(2)
	}
	if *transitKey == "" {
		f.Usage()
		os.Exit(2)
	}

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
	case "push":
		if len(args) != 1 {
			f.Usage()
			os.Exit(2)
		}
		if err := app.push(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		f.Usage()
		os.Exit(2)
	}

	os.Exit(0)
}
