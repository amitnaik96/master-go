package main

import (
	b64 "encoding/base64"
	"fmt"
)

func base64Fn() {
	data := "hello"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
}
