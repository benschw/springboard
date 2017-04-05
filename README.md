# Springboard

`springboard` is a cli utility to help get your secrets into [vault](https://www.vaultproject.io)

It leverages the [transit](https://www.vaultproject.io/docs/secrets/transit/index.html) secret backend to 
store your secrets locally in a yaml formatted secrets file and facilitates pushing
the secrets stored in this file into a specified path of the
[generic](https://www.vaultproject.io/docs/secrets/generic/index.html) secret backend.

# Usage

## configure vault for test



	docker run -d --name vault -p 8200:8200 --cap-add=IPC_LOCK -e VAULT_DEV_ROOT_TOKEN_ID=springboard vault:0.6.5

	docker run -it --rm --link vault:vault -e VAULT_TOKEN=springboard vault:0.6.5  \
		/bin/sh -c 'VAULT_ADDR=http://$VAULT_PORT_8200_TCP_ADDR:8200 vault mount transit'

	docker run -it --rm --link vault:vault -e VAULT_TOKEN=springboard vault:0.6.5  \
		/bin/sh -c 'VAULT_ADDR=http://$VAULT_PORT_8200_TCP_ADDR:8200 vault write -f transit/keys/my-key'


## use the app

	export VAULT_TOKEN=springboard
	export VAULT_ADDR=http://localhost:8200 

	./springboard set -s ./test.yml -t my-key foo "hello world"
	./springboard set -s ./test.yml -t my-key bar "hello galaxy"
	./springboard get -s ./test.yml -t my-key foo
	./springboard push -s ./test.yml -t my-key secret/my-secrets

	docker run -it --rm --link vault:vault -e VAULT_TOKEN=springboard vault:0.6.5  \
		/bin/sh -c 'VAULT_ADDR=http://$VAULT_PORT_8200_TCP_ADDR:8200 vault read secret/my-secrets'
	Key                     Value
	---                     -----
	refresh_interval        768h0m0s
	bar                     hello galaxy
	foo                     hello world

