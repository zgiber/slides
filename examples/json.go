package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type cat struct {
	Legs     int
	MaxSpeed int
	Name     string
}

// pretty print a json string of a struct using json.Indent()
func (self *someAnimal) String() string {
	s, _ := json.Marshal(self)
	var pretty bytes.Buffer
	json.Indent(&pretty, s, "", "  ")
	return string(pretty.Bytes())
}

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

// MarshalJSON enables the cat to be marhsalled as
// an animal right away (instead of just being a serial cat)
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
// stop1 OMIT
func newCat() *cat {
	c := cat{
		Legs:     4,
		MaxSpeed: 48,
		Name:     "Merin",
	}
	return &c
}

// start1 OMIT

// playground... to be structured into smaller pieces for the slides
func main() {
	// encode a cat into an animal
	// stop OMIT
	jsonCat, err := json.Marshal(newCat())
	if err != nil {
		fmt.Println(err)
	}
	// start OMIT

	// encode a shark into an animal
	jsonShark, err := json.Marshal(newShark())
	if err != nil {
		fmt.Println(err)
	}

	// see if it worked
	fmt.Println(string(jsonCat))
	fmt.Println(string(jsonShark))

	// unmarshal them into someAnimal
	cat := someAnimal{}
	shark := someAnimal{}
	err = json.Unmarshal(jsonCat, &cat)
	err = json.Unmarshal(jsonShark, &shark)

	// see if that worked (should print properties in compact json)
	fmt.Printf(
		"cat properties: %s \nshark properties: %s \n",
		string(cat.Properties),
		string(shark.Properties),
	)

	// create a []byte block of encoded animals using json.Marshal()
	blockOfAnimals := []byte{}
	for _, animal := range []interface{}{newCat(), newShark(), newCat(), newShark(), newCat()} {
		nextAnimal, err := json.Marshal(animal)
		if err != nil {
			fmt.Println(err)
		}
		blockOfAnimals = append(blockOfAnimals, nextAnimal...)
	}

	// TODO: create a []byte block of encoded animals using encoder
	// code here

	// decode a []byte block of animals using decoder:
	dec := json.NewDecoder(bytes.NewReader(blockOfAnimals))

	for {
		a := &someAnimal{}
		err := dec.Decode(a)
		if err != nil {
			if err == io.EOF {
				break
			}
			switch err.(type) {
			case *json.UnmarshalFieldError:
				// handle error
			case *json.InvalidUnmarshalError:
				// handle error
			case *json.UnmarshalTypeError:
				// handle error
			case *json.SyntaxError:
				// handle error
			}
		}
		fmt.Print(a)
	}

	// TODO: decode cats and sharks intelligently using someAnimal's "Kind" field and json.RawMessage

	// TODO: show some flags
}
