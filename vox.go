package vox

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Vox struct {
	out      io.Writer
	mu       sync.Mutex
	buf      []byte
	in       io.Reader
	progress *progress
}

var v *Vox

func init() {
	v = New()
}

// New - creates a new Vox instance. This can be used as an alternative to the
// singletone instance. If multiple Vox instances are needed.
func New() *Vox {
	v := &Vox{
		out: os.Stdout,
		in:  os.Stdin,
	}
	return v
}

// Sprintc - Creates a string in a given color. The color code prepends the
// string and a reset code is appended to it.
func Sprintc(c Color, args ...interface{}) string {
	str := fmt.Sprint(args...)
	return fmt.Sprint(c, str, ResetColor)
}

// SetOutput - Sets the output for a print actions.  By default it is Stdout.
func SetOutput(w io.Writer) { v.SetOutput(w) }
func (v *Vox) SetOutput(w io.Writer) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.out = w
}

// Output - Prints a string to the output stream.
func Output(s string) error { return v.Output(s) }
func (v *Vox) Output(s string) error {
	v.buf = v.buf[:0]
	v.buf = append(v.buf, s...)
	v.out.Write(v.buf)
	return nil
}

// Printf - Prints a formatted string using a template and as series of
// variables.
func Printf(format string, s ...interface{}) { v.Printf(format, s...) }
func (v *Vox) Printf(format string, s ...interface{}) {
	v.Output(fmt.Sprintf(format, s...))
}

// Print - Prints a number of variables.
func Print(s ...interface{})          { v.Print(s...) }
func (v *Vox) Print(s ...interface{}) { v.Output(fmt.Sprint(s...)) }

// Println - Prints a number of tokens ending with a new line.
func Println(s ...interface{}) { v.Print(s...) }
func (v *Vox) Println(s ...interface{}) {
	v.Output(fmt.Sprint(append(s, "\n")...))
}

// Printlnc - Prints a number of tokens followed by a new line. This output is
// also wrapped in a color code and a reset.
func Printlnc(c Color, s ...interface{}) { v.Printlnc(c, s...) }
func (v *Vox) Printlnc(c Color, s ...interface{}) {
	outStr := fmt.Sprint(s...)
	v.Println(c, outStr, ResetColor)
}

// Prompt - Gets input from the input stream. By default Stdin. If an empty
// string is sent the default value will be returned.
func Prompt(name, defaultVal string) string { return v.Prompt(name, defaultVal) }
func (v *Vox) Prompt(name, defaultValue string) string {
	reader := bufio.NewReader(v.in)
	Printf("%s%s [%s]: %s", Yellow, name, defaultValue, ResetColor)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	println(input)
	if input == "" {
		return defaultValue
	}
	return input
}

// PrintProperty - Prints a property name and value. The value will be right
// aligned.
func PrintProperty(name, value string) { v.PrintProperty(name, value) }
func (v *Vox) PrintProperty(name, value string) {
	totalLength := len(name) + len(value)
	if totalLength > 60 {
		Println(Yellow, name, "\n", White, value, ResetColor)
	} else {
		spaces := strings.Repeat(" ", 60-totalLength)
		Println(Yellow, name, spaces, White, value, ResetColor)
	}
}

// PrintResult - Prints a name and a result message. If an error is passed it
// will result in a failure message ex:
// Item Description                     [FAIL]
// If nil is passed as the second argument it will result in a success
// output:
// Item Description                     [OK]
// The status code will also be right aligned.
func PrintResult(desc string, err error) { v.PrintResult(desc, err) }
func (v *Vox) PrintResult(desc string, err error) {
	resultColor := Red
	resultText := "FAIL"
	if err == nil {
		resultColor = Green
		resultText = "OK"
	}
	desc += strings.Repeat(" ", 60-len(desc))
	Println(White, desc, Yellow, "[", resultColor, resultText, Yellow, "]",
		ResetColor)
}
