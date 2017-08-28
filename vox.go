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
	in       *os.File
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

func SetInput(in *os.File) { v.SetInput(in) }
func (v *Vox) SetInput(in *os.File) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.in = in
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
		v.Println(Yellow, name, "\n", White, value, ResetColor)
	} else {
		spaces := strings.Repeat(" ", 60-totalLength)
		v.Println(Yellow, name, spaces, White, value, ResetColor)
	}
}

// PrintResult - Prints a name and a result message. If an error is passed it
// will result in a failure message ex. If nil is passed as the second argument
// it will result in a success. The status code will also be right aligned and
// color coded based on the result.
func PrintResult(desc string, err error) { v.PrintResult(desc, err) }
func (v *Vox) PrintResult(desc string, err error) {
	resultColor := Red
	resultText := "FAIL"
	if err == nil {
		resultColor = Green
		resultText = "OK"
	}
	desc += strings.Repeat(" ", 60-len(desc))
	v.Println(White, desc, Yellow, "[", resultColor, resultText, Yellow, "]",
		ResetColor)
}

// Errorf - Print error output. Console output is colored red.
func Errorf(format string, args ...interface{}) { v.Errorf(format, args...) }
func (v *Vox) Errorf(format string, args ...interface{}) {
	v.Error(fmt.Sprintf(format, args...))
}

// Error - Print output as an error. Console output is colored red.
func Error(args ...interface{}) { v.Error(args...) }
func (v *Vox) Error(args ...interface{}) {
	v.Printlnc(Red, fmt.Sprint(args...))
}

// Info - Print an info output. Console output is colored white.
func Infof(format string, args ...interface{}) { v.Infof(format, args...) }
func (v *Vox) Infof(format string, args ...interface{}) {
	v.Info(fmt.Sprintf(format, args...))
}

// Info - Print an info output. Console output is colored white.
func Info(args ...interface{}) { v.Info(args...) }
func (v *Vox) Info(args ...interface{}) {
	v.Printlnc(White, fmt.Sprint(args...))
}

// Alert - Print an info output. Console output is colored yellow.
func Alertf(format string, args ...interface{}) { v.Alertf(format, args...) }
func (v *Vox) Alertf(format string, args ...interface{}) {
	v.Alertf(fmt.Sprintf(format, args...))
}

// Alert - Print an info output. Console output is colored yellow.
func Alert(args ...interface{}) { v.Alert(args...) }
func (v *Vox) Alert(args ...interface{}) {
	v.Printlnc(Yellow, fmt.Sprint(args...))
}

// Debug - Print an debug output. Debug output is not colored.
func Debugf(format string, args ...interface{}) { v.Debugf(format, args...) }
func (v *Vox) Debugf(format string, args ...interface{}) {
	v.Debug(fmt.Sprintf(format, args...))
}

// Debug - Print an debug output. Debug output is not colored.
func Debug(args ...interface{}) { v.Debug(args...) }
func (v *Vox) Debug(args ...interface{}) {
	v.Println(fmt.Sprint(args...))
}

// Fatal - Prints an error message and then exits the application.
func Fatal(args ...interface{}) { v.Fatal(args...) }
func (v *Vox) Fatal(args ...interface{}) {
	v.Error(args...)
	os.Exit(-1)
}

// Fatalf - Prints an error message and then exits the application.
func Fatalf(format string, args ...interface{}) { v.Fatalf(format, args...) }
func (v *Vox) Fatalf(format string, args ...interface{}) {
	v.Fatal(fmt.Sprintf(format, args...))
}
