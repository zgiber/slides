package main

import (
	"encoding/json"
	"fmt"
)

// stop OMIT
type cat struct {
	Legs     int
	MaxSpeed int
	Name     string
}

func main() {
	kittycat := &cat{4, 48, "Merlin"}
	jsonCat, err := json.Marshal(kittycat)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonCat))
}

// start OMIT
