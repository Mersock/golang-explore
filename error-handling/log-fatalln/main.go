package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln("err", err)
	}
}

func foo() {
	fmt.Println("defered functions don't run")
}
