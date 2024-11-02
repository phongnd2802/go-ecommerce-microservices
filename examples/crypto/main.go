package main

import (
	"fmt"

	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/crypto"
)


func main() {
	value := "duyphong02802@gmail.com"
	hashedValue := crypto.GetHash(value)
	fmt.Println(hashedValue)
}