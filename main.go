// Make json pretty
// cat foo.json | jaypwn
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

	jsbytes, err := json.MarshalIndent(out, "", "\t")
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	if len(jsbytes) == 0 {
		// didnt work, try plain obj
		out = map[string]interface{}{}
		goto try_again
	}

	io.Copy(os.Stdout, colorful(jsbytes))
	fmt.Println()
}

const (
	color_number   = "\x1b[1;37;40m"
	color_keyStr   = "\x1b[1;34;40m"
	color_valueStr = "\x1b[5;34;40m"
	color_operator = "\x1b[1;30;40m"
)

var (
	buf bytes.Buffer
)

// sure why not
func colorful(bs []byte) *bytes.Buffer {

	// toggle to color keys different from values
	isKey := true

	for _, b := range bs {

		switch {
		case b == '"':
			// highlight string values
			if isKey {
				stringColor(color_keyStr)
			} else {
				stringColor(color_valueStr)
			}
		case b >= '0' && b <= '9':
			// hilight numbers (if not in string)
			if !stringToggle {
				startColor(color_number)
				buf.WriteByte(b)
				stopColor()
			} else {
				buf.WriteByte(b)
			}
		case b == '{' || b == '}' || b == '[' || b == ']' || b == ',':
			// hilight operators
			startColor(color_operator)
			buf.WriteByte(b)
			stopColor()
		case b == ':':
			isKey = false
			buf.WriteByte(b)
		case b == '\n':
			isKey = true
			buf.WriteByte(b)
		default:
			buf.WriteByte(b)
		}
	}

	stopColor() // make sure we always stop color just in case
	return &buf
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
		buf.WriteByte('"')
	} else {
		// end string color and write closing "
		buf.WriteByte('"')
		stopColor()
	}
}

func startColor(color string) {
	buf.WriteString(color)
}

func stopColor() {
	buf.WriteString("\x1b[0m")
}
