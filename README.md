# vault-cub


	./vault-cub set -secrets foo -vault http://$HORDE_IP:8200 -path cubbyhole/foo -key userb -value passb
	./vault-cub get -secrets foo -vault http://$HORDE_IP:8200 -path cubbyhole/foo

	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 vault write cubbyhole/foo \
		a=b
	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 vault read cubbyhole/foo


# example stuff...

	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 vault mount transit

	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 vault write -f transit/keys/cub


	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 ./vault-cub set -path asd -key asd -value hellooo
	vault:v1:m0amrAD4HUPQ1qQqg+YqiMiSlZb3Jt2Kx1YJ6OhFgJT0obk=

	VAULT_TOKEN=horde VAULT_ADDR=http://$HORDE_IP:8200 ./vault-cub get -path asd -value vault:v1:uOF/dZRqijfaDBdcYmVIjmZ5EciQ2VskD8VhAjgrJ18ZJMQ=
	hellooo

