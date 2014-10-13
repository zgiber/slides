package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"time"
	// "encoding/json"
	"fmt"
	"os"
)

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

func newShark() *shark {
	s := shark{
		Fins:     6,
		MaxDepth: 3700,
		Name:     "Joe",
	}
	return &s
}

func newCat() *cat {
	c := cat{
		Legs:     4,
		MaxSpeed: 48,
		Name:     "Merin",
	}
	return &c
}

func prettyPrint(s []byte) string {
	var pretty bytes.Buffer
	json.Indent(&pretty, s, "", "  ")
	return string(pretty.Bytes())
}

func (self *cat) String() string {
	s, _ := json.Marshal(self)
	return prettyPrint(s)
}

func (self *shark) String() string {
	s, _ := json.Marshal(self)
	return prettyPrint(s)
}

func randomBool() bool {
	b := bufio.NewReader(rand.Reader)
	num, _ := binary.ReadVarint(b)
	return 0 < num&1
}

// stop OMIT
func newCatOrShark() interface{} {
	if randomBool() {
		return newCat()
	}
	return newShark()
}

func save1000animals() {
	file, _ := os.Create("animals.json")
	defer file.Close()
	enc := json.NewEncoder(file)
	for i := 0; i < 1000; i++ {
		enc.Encode(newCatOrShark())
	}
}

// start OMIT

func main() {
	//saves cats and sharks (randomly) as "animal" in ./animals
	t := time.Now()
	save1000animals()
	fmt.Println("1000 animals were saved in", time.Since(t))
}
