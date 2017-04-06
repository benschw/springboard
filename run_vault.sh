#!/bin/bash


rm -f vault

wget http://dl.fligl.io/vault.gz
gunzip vault.gz
chmod +x vault
ls -alh vault
./vault server -dev -dev-root-token-id=horde &
sleep 3
./vault mount transit
./vault write -f transit/keys/my-key
ls -alh vault
