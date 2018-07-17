package vox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

// PrintJSON - Prints a byte array contianing JSON content. This output will be
// color coded and syntax highlighted. It is also reformatted with indentation.
func PrintJSON(contentBytes []byte) { v.PrintJSON(contentBytes) }
func (v *Vox) PrintJSON(contentBytes []byte) {
	var (
		out     bytes.Buffer
		content string
		err     error
	)
	err = json.Indent(&out, contentBytes, "", "  ")
	if err == nil {
		content = string(out.Bytes())
	} else {
		println(err.Error())
	}

	v.outputPlain(content)

	re := regexp.MustCompile(`([\[\]\{\}]{1})`)
	content = re.ReplaceAllString(content, Sprintc(Green, "$1"))

	// String values
	re = regexp.MustCompile(`(\s*?\")([^:]*?)(\"\s*?,?\n)`)
	content = re.ReplaceAllString(content, fmt.Sprintf("$1%s$2%s$3", Blue, ResetColor))

	re = regexp.MustCompile(`(\:\s*[true|false]+\s*[,\n])`)
	content = re.ReplaceAllString(content, Sprintc(Magenta, "$1"))

	re = regexp.MustCompile(`(\:\s*[0-9]+\s*[,\n])`)
	content = re.ReplaceAllString(content, Sprintc(Yellow, "$1"))

	re = regexp.MustCompile(`(\:\s*null\s*[,\n])`)
	content = re.ReplaceAllString(content, Sprintc(Red, "$1"))

	v.output(content)
}
