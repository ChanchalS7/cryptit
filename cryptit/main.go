package main

import (
	"fmt"

	"github.com/ChanchalS7/cryptit/decrypt"
	"github.com/ChanchalS7/cryptit/encrypt"
)

func main(){
	encryptedStr:=encrypt.Nimbus("Kodekloud")
	// fmt.Println(encrypt.Nimbus("Kodekloud"))
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}