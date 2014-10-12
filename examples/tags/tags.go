package main

import (
	"encoding/json"
	"fmt"
)

// stop OMIT
type cat struct {
	Legs      int    `json:"limbs,omitempty"`
	MaxSpeed  int    `json:"speed,omitempty"`
	Name      string `json:"name"`
	LatinName string `json:"-"`
}

// start OMIT

func main() {
	// stop1 OMIT
	cat1 := &cat{Legs: 4, MaxSpeed: 48, Name: "Merlin"}
	// start1 OMIT
	jsonCat, _ := json.Marshal(cat1)
	fmt.Println(string(jsonCat))
}
