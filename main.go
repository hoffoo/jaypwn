// Make json pretty
// cat foo.json | jsonfu
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// to json
	dec := json.NewDecoder(os.Stdin)
	out := map[string]interface{}{}
	dec.Decode(&out)

	pretty, err := json.MarshalIndent(out, "", "\t")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	os.Stdout.Write(pretty)
}
