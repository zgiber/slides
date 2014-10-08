package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// stop OMIT
func main() {
	cat := `
{
  "Kind": "cat",
  "Properties": {
    "Legs": 4,
    "MaxSpeed": 48,
    "Name": "Merin"
  }
}
`
	b := bytes.Buffer{}
	json.Compact(&b, []byte(cat))
	fmt.Println(string(b.Bytes()))
}

// start OMIT
