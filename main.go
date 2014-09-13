// Make json pretty
// cat foo.json | jsonfu
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// either an object or array
	var out interface{}

	// to json
	decoder := json.NewDecoder(os.Stdin)

	out = []map[string]interface{}{}

	// first try array, if it fails (or empty) try plain obj
try_again:
	decoder.Decode(&out)

	pretty, err := json.MarshalIndent(out, "", "\t")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	if len(pretty) == 0 {
		// didnt work, try plain obj
		out = map[string]interface{}{}
		goto try_again
	}

	os.Stdout.Write(pretty)
}
