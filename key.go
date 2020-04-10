package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
)

const pubkey = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlgbfxMniLiDMRYhRYY0f
fOLEACXSCWmX0/rSL+qib/3cbZAMOEknXUudich4ZuCHulZ9ApaPx/7u+x5jQSj4
aiZXJE+S+LecUqwbq1CSfByLPViyYu2xt2I0tqYdsQK6KmQs2Gl00UP/yxrHtcEz
NaZO8Z7bdL1AY3eW6oPjWeORK91FAEONbnCvXmPoGa/4+AUWr6FmMrjFiG8yM72K
eUvfzyWtZYNeFxJ+2UmqTco1oEdGmwJJYKgPAg4mRXOPBs1Il6W9+bwomUed/Rxd
GHuNPy4b9BOgSyFFoEHQJ2eL+W9IMpWegwV7VxXc37WlHQxoZ1886gO+u3hxvo++
+v0ami3JT1BZriTYdjSydktyUARQQzDaxAsYwUMTs/G++yiF3jt+J43pKvZ+ZSTP
+vXAKd+acbsUmH6WIxsu915BVPcnMgyeUWOK6NojiW4Z4BEuCWVKfqMKRU+LypFN
Hqpd3wxT26jnykJOm0a2xloXlmjS9x/LcHd6onN6I6wdPz8zSAU6lr0T2kWgPY+l
u0Ral9lpafe/Rq6GjPIvrlWNy2hjJhJ1FtzMCgySCs+XEqjFbM2GEOSK4M/NGY9+
zzkNgL4B0HpMHgRNeRfx0q+LuZtuHvNEDxmp/OvvfRqQGo5qqDhojm3rRi5qbsLa
k3siF46a7ml6ONtAD/Eib1kCAwEAAQ==
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
