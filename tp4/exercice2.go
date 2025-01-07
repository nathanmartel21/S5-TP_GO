package main

import "fmt"

func f(x float64) float64 {

	return x*x + 3

}

func main() {

	var a, b int32

	fmt.Println("Donnez un nombre a : ")
	fmt.Scan(&a)

	fmt.Println("Donnez un nombre b : ")
	fmt.Scan(&b)

	var som float32

	for x := 1; x < 4; x++ {

		som += f(x)

	}

	var h = (b - a) / 4

	som += h

	fmt.Println(som)

}
