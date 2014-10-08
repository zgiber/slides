package main

import (
	"encoding/json"
	"fmt"
)

// stop OMIT
type shark struct {
	Fins     int
	MaxDepth int
	Name     string
}

func main() {
	jsonShark := []byte(`{"Fins":6,"MaxDepth": 3700,"Name": "Joe"}`)
	s := &shark{}
	err := json.Unmarshal(jsonShark, s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s the shark has %v fins so he can swim %vm deep.\n",
		s.Name, s.Fins, s.MaxDepth)
}

// start OMIT
