package main

import "fmt"
import "encoding/json"

type person struct {
	Name string
}

// stop OMIT
func main() {
	p1 := &person{}
	p2 := person{}
	input := []byte(`{"name":"John Doe"}`)
	json.Unmarshal(input, p1)
	json.Unmarshal(input, p2)
	fmt.Println("person1:", p1.Name, "person2:", p2.Name)
}

// start OMIT
