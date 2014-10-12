package main

import "fmt"
import "encoding/json"

type unknownJSON struct {
	Decoded interface{}
}

// stop OMIT
func get(result interface{}, key string) interface{} {
	assertedMap, ok := result.(map[string]interface{})
	// start OMIT
	if !ok {
		return "invalid key"
	}
	// stop1 OMIT
	if v, exists := assertedMap[key]; exists {
		return v
	}
	// start1 OMIT
	return ""
}

//stop2 OMIT
func main() {
	something := []byte(`{"thing":{"insideThing":{"insideInsideThing":{"value":42}}}}`)
	var decoded interface{}
	_ = json.Unmarshal(something, &decoded)
	//	can't use it right away: OMIT
	//	fmt.Println(decoded["thing"], err) OMIT
	fmt.Println(
		get(get(get(get(decoded, "thing"), "insideThing"), "insideInsideThing"), "value"))
}

// start2 OMIT
