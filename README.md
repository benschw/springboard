# vault-cub


	./vault-cub set -secrets foo -vault http://$HORDE_IP:8200 -path cubbyhole/foo -key userb -value passb
	./vault-cub get -secrets foo -vault http://$HORDE_IP:8200 -path cubbyhole/foo

	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 vault write cubbyhole/foo \
		a=b
	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 vault read cubbyhole/foo


# example stuff...

	export VAULT_TOKEN=horde
	export VAULT_ADDR=http://$HORDE_IP:8200 

	vault mount transit
	vault write -f transit/keys/cub

	touch test.yml

	./vault-cub set -s ./test.yml -t cub foo "hello world"
	./vault-cub get -s ./test.yml -t cub foo
