package main

import (
	"fmt"
	"log"
)

type norgatMathError struct {
	lat  string
	long string
	err  error
}

func (n norgatMathError) Error() string {
	return fmt.Sprintf("a norgat math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.22)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root negative number: %v", f)
		return 0, norgatMathError{"12.22 N", "33.33 W", nme}
	}
	return 42, nil
}
