package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

const pubkey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAx6WcU5phXluhQHvzna0S
R0q/7z2Fh9qazecvoaiYONS8QHrWDMSNMcuHr1Ulm2yB+psd99Y0TFF5Vb5iBr1m
UJ+1uJFzzYIVyb1ioNVv4WJmRQKCUVlL47E9FE1/nfu1qtez3ZPGwjHwxgiXwbzE
fKgAiwN1f8sgCwVnhInRw+dquJ2y+BoIcA/wUBJ0zkHe/6ud9I+XWlqDOYMIXEXD
1IhxFM9aiZj3S/Uy1+OmgxOIVJbtOkn6PYJZQGBNOuPztEhMbuRa2/6zEguazmXo
ElLphUQG3xHQ+GedcZCofqZhbPCaxZQJ5ioBcZd1Nr73r0JWJ+VLG2iODCTRoHBr
DQIDAQAB
-----END PUBLIC KEY-----
`

var PublicKey *rsa.PublicKey

func init() {
	var err error
	block, _ := pem.Decode([]byte(pubkey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	PublicKey = pub.(*rsa.PublicKey)
}
