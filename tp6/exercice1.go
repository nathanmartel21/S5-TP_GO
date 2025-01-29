package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Interface Mouvement
type Mouvement interface {
	Vitesse() int
	Nom() string
}

// Structures des véhicules
type Voiture struct{}
type Moto struct{}
type Velo struct{}

func (v Voiture) Vitesse() int { return rand.Intn(151) + 50 } // 50 - 200 km/h
func (v Moto) Vitesse() int    { return rand.Intn(121) + 30 } // 30 - 150 km/h
func (v Velo) Vitesse() int    { return rand.Intn(31) + 5 }   // 5 - 35 km/h

func (v Voiture) Nom() string { return "Voiture" }
func (v Moto) Nom() string    { return "Moto" }
func (v Velo) Nom() string    { return "Vélo" }

// Fonction pour initialiser un tableau de véhicules
func genererVehicules(n int) []Mouvement {
	rand.Seed(time.Now().UnixNano())
	vehicules := []Mouvement{}

	for i := 0; i < n; i++ {
		switch rand.Intn(3) {
		case 0:
			vehicules = append(vehicules, Voiture{})
		case 1:
			vehicules = append(vehicules, Moto{})
		case 2:
			vehicules = append(vehicules, Velo{})
		}
	}

	return vehicules
}

// Fonction pour afficher la vitesse des véhicules
func afficherVitesses(vehicules []Mouvement) {
	var maxVitesse int
	var vehiculeMax string

	fmt.Println("Vitesse des véhicules enregistrées :")
	for _, v := range vehicules {
		vitesse := v.Vitesse()
		fmt.Printf("%s : %d km/h\n", v.Nom(), vitesse)
		if vitesse > maxVitesse {
			maxVitesse = vitesse
			vehiculeMax = v.Nom()
		}
	}
	fmt.Printf("Vitesse maximale : %d km/h (%s)\n", maxVitesse, vehiculeMax)
}

// Fonction pour sauvegarder les données dans un fichier YAML
func sauvegarderYAML(vehicules []Mouvement, filename string) {
	data := []map[string]interface{}{}

	for _, v := range vehicules {
		data = append(data, map[string]interface{}{
			"type":    v.Nom(),
			"vitesse": v.Vitesse(),
		})
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en YAML :", err)
		return
	}

	err = os.WriteFile(filename, yamlData, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}

	fmt.Println("Données enregistrées dans", filename)
}

func main() {
	vehicules := genererVehicules(10) // Génère 10 véhicules
	afficherVitesses(vehicules)
	sauvegarderYAML(vehicules, "vehicules.yaml")
}

/*Interface Mouvement avec Vitesse() et Nom().
Structures Voiture, Moto, Velo qui définissent leur propre vitesse aléatoire.
Fonction genererVehicules(n int) : crée un tableau de n véhicules aléatoires.
Fonction afficherVitesses(vehicules) : affiche les vitesses et la vitesse max.
Fonction sauvegarderYAML(vehicules, filename) : stocke les données dans un fichier YAML.
main() :
Génère 10 véhicules.
Affiche leurs vitesses.
Sauvegarde les résultats en YAML.*/
