package main

import (
	"fmt"

	"math/big"
	// tuples "github.com/riavalon/ray_tracer/tuples"
)

func main() {
	fmt.Println("Dividing 12.6 by 2")

	x, _ := new(big.Float).SetPrec(200).SetString("4.0")
	y, _ := new(big.Float).SetPrec(200).SetString("2")
	result := new(big.Float).Quo(x, y)
	fmt.Println(result.String())
}
