package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type secretAgent struct {
	Name        string
	FakeName    string
	Nationality string
	Password    string
}

// switch names on the wire
func (self *secretAgent) MarshalJSON() ([]byte, error) {
	publicInfo := struct {
		Name        string
		FakeName    string
		Nationality string
	}{self.FakeName, self.Name, randomPick()}
	return json.Marshal(&publicInfo)
}

func randomPick() string {
	choices := []string{
		"italian",
		"english",
		"russian",
		"german",
		"american",
	}
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Int()%5]
}

func main() {
	agent := &secretAgent{"Zoltan", "John", "hungarian", "verystrongpwd"}
	b, err := json.Marshal(agent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
