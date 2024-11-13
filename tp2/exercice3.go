package main

import "fmt"

func main() {

	var nombre_homme int8 = 75
	var nombre_femme int8 = 46
	var nombre = int16(nombre_homme) + int16(nombre_femme)

	var salaire_mensuel int16 = 1982
	var salaire_mensuel_homme int16 = 2040

	var salaire_mensuel_femme float32 = 0

	//1982 = (2040*75 + salaire_mensuel_femme*46)/121
	//1982 * 121 = 2040*75 + salaire_mensuel_femme*46
	//1982 * 121 - (2040*75) = salaire_mensuel_femme*46

	salaire_mensuel_femme = ((float32(salaire_mensuel)*float32(nombre) - (float32(salaire_mensuel_homme) * float32(nombre_homme))) / float32(nombre_femme))

	fmt.Println(int32(salaire_mensuel_femme))

}
