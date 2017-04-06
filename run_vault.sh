#!/bin/bash


rm -f vault

wget http://dl.fligl.io/vault -O /vault
chmod 755 /vault
/vault server -dev -dev-root-token-id=horde &
sleep 3
/vault mount transit
/vault write -f transit/keys/my-key
