package main

import (
	"github.com/khalidMujahid/src"
)

func main() {

	if err := src.Run(); err != nil {
		panic(err)
	}
}
