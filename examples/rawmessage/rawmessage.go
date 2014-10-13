package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// stop OMIT
type someAnimal struct {
	Kind       string
	Properties json.RawMessage
}

func main() {
	url := `http://127.0.0.1:3999/animals.json`
	resp, _ := http.Get(url)
	dec := json.NewDecoder(resp.Body)

	for i := 0; i < 10; i++ {
		a := &someAnimal{}
		dec.Decode(a)
		fmt.Println(a.Kind, a.Properties)
	}
	// start OMIT
	resp.Body.Close()
}
