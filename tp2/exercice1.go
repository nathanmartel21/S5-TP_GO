package main

import "fmt"

func main() {

	var taille_enfant1 int16 = 116
	var taille_enfant2 int16 = 212
	var taille_enfant3 int16 = 114
	var taille_enfant4 int16 = 128

	var nombre_enfant int16 = 4

	var moyenne float32 = (float32(taille_enfant1) + float32(taille_enfant2) + float32(taille_enfant3) + float32(taille_enfant4)) / float32(nombre_enfant)

	fmt.Println(moyenne)

}
