package main

import "fmt"

// Définition de la structure Sheep
type Sheep struct {
	name   string
	age    int
	weight float64
}

// Fonction pour créer un nouveau mouton
func NewSheep(name string, age int, weight float64) Sheep {
	return Sheep{name: name, age: age, weight: weight}
}

// Méthode pour afficher les informations d'un mouton
func (s Sheep) PrintSheep() {
	fmt.Printf("Nom: %s, Âge: %d ans, Poids: %.2f kg\n", s.name, s.age, s.weight)
}

// Définition de la ferme (map qui stocke les moutons)
type Farm map[string]Sheep

// Fonction pour ajouter un mouton à la ferme
func (f Farm) AddSheep(s Sheep) {
	f[s.name] = s
}

// Fonction pour supprimer un mouton de la ferme
func (f Farm) RemoveSheep(name string) {
	delete(f, name)
}

// Fonction pour obtenir un mouton par son nom
func (f Farm) GetSheep(name string) (Sheep, bool) {
	sheep, exists := f[name]
	return sheep, exists
}

// Fonction pour compter le nombre de moutons dans la ferme
func (f Farm) CountSheeps() int {
	return len(f)
}

func main() {
	// Création de la ferme
	farm := Farm{}

	// Ajout de moutons
	farm.AddSheep(NewSheep("Dolly", 3, 75.5))
	farm.AddSheep(NewSheep("Shaun", 2, 50.2))
	farm.AddSheep(NewSheep("Béééla", 4, 60.7))

	// Affichage des moutons
	fmt.Println("Moutons dans la ferme :")
	for _, sheep := range farm {
		sheep.PrintSheep()
	}

	// Nombre de moutons
	fmt.Printf("Nombre total de moutons : %d\n", farm.CountSheeps())

	// Récupération d'un mouton
	if sheep, found := farm.GetSheep("Shaun"); found {
		fmt.Println("Mouton récupéré :")
		sheep.PrintSheep()
	} else {
		fmt.Println("Shaun n'est pas dans la ferme.")
	}

	// Suppression d'un mouton
	farm.RemoveSheep("Dolly")
	fmt.Printf("Après suppression, il reste %d moutons.\n", farm.CountSheeps())
}

/*
Correction de NewSheep() : il utilise bien les paramètres d'entrée.
Ajout de la map[string]Sheep nommée Farm pour stocker les moutons.
Ajout de méthodes sur Farm :
AddSheep(s Sheep): Ajoute un mouton.
RemoveSheep(name string): Supprime un mouton.
GetSheep(name string): Récupère un mouton.
CountSheeps(): Compte le nombre de moutons.
Dans main() :
Création d'une ferme.
Ajout de moutons.
Affichage des moutons.
Recherche et suppression d'un mouton.
Affichage du nombre de moutons restants.
*/
