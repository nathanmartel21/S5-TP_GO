package main

import (
	"fmt"
	"math"
)

func points() ([]int32, []int32) {
	tab_abscisse := []int32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	tab_ordonnees := []int32{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	// x := make([]int, n)) => tableau de taille . COmme un malloc.
	return tab_abscisse, tab_ordonnees
}

func barycentre(tab_abscisse []int32, tab_ordonnees []int32) (float64, float64) {
	var sommeX, sommeY int32
	n := len(tab_abscisse)

	for i := 0; i < n; i++ {
		sommeX += tab_abscisse[i]
		sommeY += tab_ordonnees[i]
	}

	return float64(sommeX) / float64(n), float64(sommeY) / float64(n)
}

func distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func main() {
	tab_abscisse, tab_ordonnees := points()
	x, y := barycentre(tab_abscisse, tab_ordonnees)

	var minDistance float64 = math.MaxFloat64
	var index int

	for j := 0; j < len(tab_abscisse); j++ {
		dist := distance(float64(tab_abscisse[j]), float64(tab_ordonnees[j]), x, y)

		if dist < minDistance {
			minDistance = dist
			index = j
		}
	}
	fmt.Printf("Barycentre : (%f, %f)\n", x, y)
	fmt.Printf("Le point le plus proche du barycentre est : (%d, %d)\n", tab_abscisse[index], tab_ordonnees[index])
}
