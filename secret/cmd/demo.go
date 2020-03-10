package main

import (
	"fmt"

	"github.com/kazijawad/Gophercises/secret"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")
	err := v.Set("demo-key1", "123 some crazy value")
	if err != nil {
		panic(err)
	}
	err = v.Set("demo-key2", "456 some crazy value")
	if err != nil {
		panic(err)
	}
	err = v.Set("demo-key3", "789 some crazy value")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo-key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain 1:", plain)
	plain, err = v.Get("demo-key2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain 2:", plain)
	plain, err = v.Get("demo-key3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain 3:", plain)
}
