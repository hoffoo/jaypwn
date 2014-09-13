// Make json pretty
// cat foo.json | jsonfu
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	if buf.Len() != 0 {
		io.Copy(os.Stdout, &buf)
		fmt.Println()
	}
}

const (
	number   = "\x1b[1;37;40m"
	keyStr   = "\x1b[1;34;40m"
	valueStr = "\x1b[5;34;40m"
	operator = "\x1b[1;30;40m"
)

// sure why not
func colorful(bs []byte) {

	// toggle to color keys different from values
	isKey := true

	for _, b := range bs {

		switch {
		case b == '"':
			// highlight string values
			if isKey {
				stringColor(keyStr)
			} else {
				stringColor(valueStr)
			}
		case rune(b) >= '0' && rune(b) <= '9':
			// hilight numbers (if not in string)
			if !stringToggle {
				startColor(number)
				writeByte(b)
				stopColor()
			} else {
				writeByte(b)
			}
		case b == '{' || b == '}':
			startColor(operator)
			writeByte(b)
			stopColor()
		case b == ':':
			isKey = false
			writeByte(b)
		case b == '\n':
			isKey = true
			writeByte(b)
		case b == '[' || b == ']' || b == ',':
			startColor(operator)
			writeByte(b)
			stopColor()
		default:
			writeByte(b)
		}
	}
}

var (
	stringToggle bool
)

func stringColor(color string) {

	stringToggle = !stringToggle

	// hilight strings
	if stringToggle {
		// start string color
		startColor(color)
		writeByte('"')
	} else {
		// end string color and write closing "
		writeByte('"')
		stopColor()
	}
}

func startColor(color string) {
	buf.WriteString(color)
}

func stopColor() {
	buf.WriteString("\x1b[0m")
}

var (
	buf bytes.Buffer
)

func writeByte(b byte) {
	buf.WriteByte(b)
}
