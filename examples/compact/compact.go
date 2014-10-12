package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// stop OMIT
func main() {
	animal := `
{
  "Kind": "shark",
  "Properties": {
    "Fins": 6,
    "MaxDepth": 3700,
    "Name": "Joe"
  }
}
`
	b := bytes.Buffer{}
	json.Compact(&b, []byte(animal))
	fmt.Println(string(b.Bytes()))
}

// start OMIT
