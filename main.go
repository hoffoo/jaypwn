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

	colorful(pretty)
}

// sure why not
func colorful(bs []byte) {

	var stringToggle bool

	for _, b := range bs {

		// yay gross
		var skipWrite bool

		// hilight strings
		if b == '"' {

			stringToggle = !stringToggle

			if stringToggle {
				// start string color
				fmt.Print("\x1b[35;1m")
			} else {
				// end string color and write closing "
				skipWrite = true
				fmt.Print("\"\x1b[0m")
			}
		}

		// hilight numbers
		if rune(b) >= '0' && rune(b) <= '9' {
			skipWrite = true

			fmt.Print("\x1b[34;1m")
			os.Stdout.Write([]byte{b})
			fmt.Printf("\x1b[0m")
		}

		// write byte unless already wrote
		if !skipWrite {
			os.Stdout.Write([]byte{b})
		}
	}
}
