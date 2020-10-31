package main

import (
	"fmt"

	"github.com/Mersock/golang-explore/make-go-doc/mersockacdc"
	"github.com/Mersock/golang-explore/make-go-doc/mersockdogpkg"
)

type canine struct {
	name string
	age  int
}

func main() {
	dog := canine{
		name: "Dog",
		age:  mersockdogpkg.Years(10),
	}
	fmt.Println(dog)
	fmt.Println(mersockacdc.Sum(2, 3))
}
