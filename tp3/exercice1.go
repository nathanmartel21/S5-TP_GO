package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var nombre int32

	var nombreale = rand.Intn(100) + 1

	var nb_essai int16 = 0
	var nb_essai_max int16 = 10

	fmt.Println(nombreale)

	for nb_essai < nb_essai_max {

		fmt.Println("Donnez un nombre entre 1 et 100")
		fmt.Scan(&nombre)

		if int32(nombreale) > int32(nombre) {
			fmt.Println("Nombre différent, chercher plus haut")
			nb_essai += 1
			fmt.Println("Nombre essais restant : ", 10-nb_essai)
		} else if int32(nombreale) < int32(nombre) {
			fmt.Println("Nombre différent, chercher plus bas")
			nb_essai += 1
			fmt.Println("Nombre essais restant : ", 10-nb_essai)
		} else {
			fmt.Println("Nombre trouvé, bravo !")
			break
		}
	}

}
