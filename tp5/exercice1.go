package main

import "fmt"

type Sheep struct {
	nom    string
	age    int
	weight float64
}

func NewSheep(nom string, age int, weight float64) Sheep {

	return Sheep{
		nom:    "nathan",
		age:    21,
		weight: 100,
	}
}

func (s Sheep) PrintSheep() {

	fmt.Println(s.nom, s.age, s.weight)

}
