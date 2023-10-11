package main

import (
	"fmt"

	"github.com/cespare/xxhash"
)

func main() {
	fmt.Println("hello world")

	hash := xxhash.New()
	hash.Write([]byte{1, 2, 3, 4})
	foo := hash.Sum64()

	fmt.Println(foo)
}
