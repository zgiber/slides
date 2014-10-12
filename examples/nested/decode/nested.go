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
	input := `
{
  "Recepient": {"Email": "joe@example.com","Name": "Joe"},
  "Sender": {"Email": "jimmy@example.com","Name": "Jimmy"},
  "Body": "Hey Joe!"
}`

	message := &Message{}
	json.Unmarshal([]byte(input), &message)
	fmt.Printf("From %s\nTo %s\n%s\n",
		message.Sender.Name, message.Recepient.Name, message.Body)
}

// start1 OMIT
