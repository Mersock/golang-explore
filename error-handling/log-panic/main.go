package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Panic("err", err)
		panic(err)
	}
}
