package main

import "fmt"

func un(n int) float64 {
	if n == 0 {
		return 0
	}

	u := 0.0 // car on retourne u qui est un float (fonction retourne un float)
	for i := 1; i <= n; i++ {
		u = 0.5*u + 3
	}

	return u
}

func smallind() int {
	u0 := 0.0
	n := 0
	for {
		u := 0.5*u0 + 3
		if abs(u-u0) <= 1e-4 {
			return n
		}
		u0 = u
		n++
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int

	fmt.Print("Donnez un indice n : ")
	fmt.Scanln(&n)

	resultat := un(n)
	fmt.Println("Le résultat à l'indice", n, "est", resultat)

	indicemin := smallind()

	fmt.Printf("L'indice minimal est %d\n", indicemin)
}
