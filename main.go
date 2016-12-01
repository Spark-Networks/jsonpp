package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var i interface{}
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&i); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing json: %v\n", err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(i); err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting json: %v\n", err)
		os.Exit(1)
	}
}
