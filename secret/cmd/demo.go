package main

import (
	"fmt"

	"github.com/kazijawad/Gophercises/secret"
)

func main() {
	v := secret.Memory("my-fake-key")
	err := v.Set("demo-key", "some crazy value")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo-key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
