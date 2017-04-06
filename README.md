# Springboard

`springboard` is a cli utility to help get your secrets into [vault](https://www.vaultproject.io)

It leverages the [transit](https://www.vaultproject.io/docs/secrets/transit/index.html) secret backend to 
store your secrets locally in a yaml formatted secrets file and facilitates pushing
the secrets stored in this file into a specified path of the
[generic](https://www.vaultproject.io/docs/secrets/generic/index.html) secret backend.

# Usage

## Vault Setup

	export VAULT_TOKEN=springboard
	export VAULT_ADDR=http://localhost:8200 

	vault server -dev -dev-root-token-id=springboard

	vault mount transit
	Successfully mounted 'transit' at 'transit'!

	vault write -f transit/keys/my-key
	Success! Data written to: transit/keys/cub


## Manage Secrets

	export VAULT_TOKEN=springboard
	export VAULT_ADDR=http://localhost:8200 

	./springboard set -s ./test.yml -t my-key foo "hello world"
	./springboard set -s ./test.yml -t my-key bar "hello galaxy"
	./springboard get -s ./test.yml -t my-key foo
	hello world
	./springboard push -s ./test.yml -t my-key secret/my-secrets

	vault read secret/my-secrets'
	Key                     Value
	---                     -----
	refresh_interval        768h0m0s
	bar                     hello galaxy
	foo                     hello world

