package main

import "fmt"

func main() {

	var note1 int8 = 10
	var note2 int8 = 13
	var note3 int8 = 14
	var note4 int8 = 9
	var note5 int8 = 8

	var nombre_notes int16 = 5

	var moyenne float32 = (float32(note1) + float32(note2) + float32(note3) + float32(note4) + float32(note5)) / float32(nombre_notes)

	fmt.Println("Moyenne actuelle :", moyenne)

	var moyenne_voulue int8 = 11

	var somme_note int8 = note1 + note2 + note3 + note4 + note5

	fmt.Println("Somme des notes : ", somme_note)

	var somme_voulue = 6 * moyenne_voulue

	var note_minimal = somme_voulue - somme_note

	fmt.Println("L'élève doit avoir au minimum la note de", note_minimal)

}
