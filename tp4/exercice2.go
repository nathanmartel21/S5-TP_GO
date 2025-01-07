package main

import "fmt"

func f(x float64) float64 {

	return x*x + 3

}

func main() {

	var a, b int32 = 1, 4
	var n int

	fmt.Println("Donnez un nombre de rectangles : ")
	fmt.Scan(&n)

	var h float64 = (float64(b) - float64(a)) / float64(n)

	var R float64

	for i := 0; i < n; i++ {

		var xi float64 = float64(a) + float64(i)*float64(h)
		R += f(xi)

	}

	R *= h

	fmt.Println(R)

}
