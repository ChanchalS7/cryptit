package main

import (
	"fmt"

	"github.com/pborman/uuid"
	"github.com/ChanchalS7/cryptit/encrypt"
	"github.com/ChanchalS7/cryptit/decrypt"
)

func main(){
	uuid:=uuid.NewRandom()
	fmt.Println(uuid)
	encryptedStr :=encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
}