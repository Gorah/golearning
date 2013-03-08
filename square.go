package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {

	if x == -1 {
		return 0, errors.New("Missing value!")
	} else if x <= 0 {
		return 0, errors.New("Value must be greater than 0")
	}

	z := 1.0
	var tmp float64
	for i := 1; i < 10; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		if z == tmp {
			return z, nil
		}
		tmp = z
	}
	return z, nil
}

func main() {

	we := flag.Int("sq", -1, "Give value to be square rooted")
	flag.Parse()
	num, err := Sqrt(float64(*we))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
