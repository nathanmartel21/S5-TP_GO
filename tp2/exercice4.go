package main

import (
	"fmt"
	"math"
)

func main() {

	var rayon int8

	fmt.Println("Donnez le rayon d'un cercle : ")
	fmt.Scan(&rayon)

	var perimetre float32 = (float32(2) * float32(math.Pi) * float32(rayon))
	var surface float32 = (float32(math.Pi) * float32(rayon) * float32(rayon))

	fmt.Println("Le perimètre du cercle est de : ", perimetre)
	fmt.Println("La surface du cercle est de : ", surface)

}
