package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// used by decoder/unmarshaler, when decoding arbitrary kind of animal
// keeping properties a JSON for further processing
type someAnimal struct {
	Kind       string
	Properties json.RawMessage
}

// animal struct used by encoder/marshaler, when encoding cats
// and sharks which have their own MarshalJSON methods, utilizing
// this the animal struct.
type animal struct {
	Kind       string
	Properties interface{}
}

type shark struct {
	Fins     int
	MaxDepth int
	Name     string
}

// stop OMIT
type cat struct {
	Legs     int
	MaxSpeed int
	Name     string
}

// start OMIT

// pretty print a json string of a struct using json.Indent()
// stop1 OMIT
func (self *cat) String() string {
	s, _ := json.Marshal(self)
	var pretty bytes.Buffer
	json.Indent(&pretty, s, "", "  ")
	return string(pretty.Bytes())
}

// start1 OMIT

// MarshalJSON method enables a shark to be marhsalled as
// an animal right away (instead of just being a serialized shark)

func (self *shark) MarshalJSON() ([]byte, error) {
	a := animal{
		Kind: "shark",
		Properties: shark{
			Fins:     self.Fins,
			MaxDepth: self.MaxDepth,
			Name:     self.Name,
		},
	}
	return json.Marshal(a)
}

// stop2 OMIT
func (self *cat) MarshalJSON() ([]byte, error) {
	a := animal{
		Kind: "cat",
		Properties: cat{
			Legs:     self.Legs,
			MaxSpeed: self.MaxSpeed,
			Name:     self.Name,
		},
	}
	return json.Marshal(a)
}

// start2 OMIT

// shark generator.. all of them Joes, with 6 fins
// look how different is that from a cat
func newShark() *shark {
	s := shark{
		Fins:     6,
		MaxDepth: 3700,
		Name:     "Joe",
	}
	return &s
}

// cat generator.. 4 legs, all of them wizards
// not a shark, right?
func newCat() *cat {
	c := cat{
		Legs:     4,
		MaxSpeed: 48,
		Name:     "Merin",
	}
	return &c
}

// playground... to be structured into smaller pieces for the slides
// stop1 OMIT
func main() {
	cat := newCat()
	fmt.Print(cat)
}

// start1 OMIT
