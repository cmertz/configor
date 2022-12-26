package main

import (
	"fmt"

	"github.com/cmertz/configor"
)

type config struct {
	A string
	B string
}

func main() {
	var c config

	err := configor.Resolve(&c, configor.Env())
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
}
