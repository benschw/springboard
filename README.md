[![Build Status](https://travis-ci.org/benschw/springboard.svg?branch=master)](https://travis-ci.org/benschw/springboard)
[![Downloads](https://img.shields.io/badge/download-release-blue.svg)](http://dl.fligl.io/#/springboard)

# Springboard

`springboard` is a cli utility to help get your secrets into [vault](https://www.vaultproject.io)

It leverages the [transit](https://www.vaultproject.io/docs/secrets/transit/index.html) secret backend to 
protect your secrets locally in a yaml formatted secrets file and facilitates pushing
the secrets stored in this file into a specified path of the
[generic](https://www.vaultproject.io/docs/secrets/generic/index.html) secret backend.

Since the values in your secrets file are encrypted with vault's transit backend, 
you can commit these files to source control. This allows you to be more deliberate about
tracking and publishing the sensitive data needed to run your applications.


# Usage

## Install

	wget http://dl.fligl.io/artifacts/springboard/springboard_linux_amd64.gz
	gunzip springboard_linux_amd64.gz
	chmod +x springboard_linux_amd64
	mv springboard_linux_amd64 /usr/local/bin/

## Configure Vault

_Hard coding tokens etc. is only suitable for dev. See
[Installing Vault](https://www.vaultproject.io/docs/install/index.html)
for complete install instructions_

	export VAULT_TOKEN=springboard
	export VAULT_ADDR=http://localhost:8200 

	vault server -dev -dev-root-token-id=springboard

	vault mount transit
	Successfully mounted 'transit' at 'transit'!

	vault write -f transit/keys/my-key
	Success! Data written to: transit/keys/my-key


## Managing Secrets with Springboard

	export VAULT_TOKEN=springboard
	export VAULT_ADDR=http://localhost:8200 

	springboard set -s ./test.yml -t my-key foo "hello world"
	springboard set -s ./test.yml -t my-key bar "hello galaxy"
	springboard get -s ./test.yml -t my-key foo
	hello world
	springboard push -s ./test.yml -t my-key secret/my-secrets

	vault read secret/my-secrets
	Key                     Value
	---                     -----
	refresh_interval        768h0m0s
	bar                     hello galaxy
	foo                     hello world

## Complete Usage

	Usage: springboard <subcommand> -s <secrets file> -t <transit key> [args]

	Subcommands:
	    help                display this help screen and exit
	    set <key> <value>   set/encrypt 'value' in local secrets file
	    get <key>           get/decrypt 'value' from local secrets file
	    push <path>         publish secrets in local secrets file to
                        'path' in vault generic secrets backend

	Flags:
	    -s string   secrets file
	    -t string   transit key

	Examples:
	    springboard set -s secrets.yml -t my-key user_name supersecret
	    springboard get -s secrets.yml -t my-key user_name
	    springboard push -s secrets.yml -t my-key secret/my-space

	github.com/benschw/springboard
