package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Email string
	Name  string
}

// stop OMIT
type Message struct {
	Recepient Person
	Sender    Person
	Body      string
}

// start OMIT

func prettyPrint(s []byte) string {
	var pretty bytes.Buffer
	json.Indent(&pretty, s, "", "  ")
	return string(pretty.Bytes())
}

// stop1 OMIT
func main() {
	message := &Message{
		Recepient: Person{Email: "joe@example.com", Name: "Joe"},
		Sender:    Person{Email: "jimmy@example.com", Name: "Jimmy"},
		Body:      "Hey Joe!",
	}

	jsonMessage, _ := json.Marshal(&message)
	fmt.Println(prettyPrint(jsonMessage))
}

// start1 OMIT
