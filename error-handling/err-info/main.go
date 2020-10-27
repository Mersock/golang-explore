package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	f, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(f)
	p, err := getPi(2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}

func getPi(f float64) (float64, error) {
	if f < 3.14 {
		err := fmt.Errorf("orgate math: square root of negative number: %v", f)
		return 0, err
	}
	return 3.14, nil
}
