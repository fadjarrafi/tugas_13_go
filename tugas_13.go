package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "Fadjar Irfan Rafi"

	// encode ke base64
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded : ", encodedString)

	// enkripsi menggunakan sha1
	var sha = sha1.New()
	sha.Write([]byte(data))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println("Hasil enkripsi : ", encryptedString)
}
