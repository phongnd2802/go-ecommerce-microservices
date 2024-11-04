package main

import (
	"fmt"

	"github.com/phongnd2802/go-ecommerce-microservices/pkg/utils/random"
)

func main() {
	fmt.Println(random.RandomString(10))
}