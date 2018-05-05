// Vox is a Go package designed to help make terminal/console applications more attractive.
// 
// It is a collection of small helper functions that aid in printing various pieces of information to the console.
// 
// - Various predefined and common printing tasks like printing property key/value pairs, result responses, etc.
// - Print JSON data with syntax highlighting
// - Easily print colorized output
// - Display real time progress bars for tasks
// - Easy helper functions for printing various types of messages: Alerts, errors, debug messages, etc.
//  - Control the output and input streams to help during application testing.
//
//  Printing to the screen
// 
// There are a number of output functions to print data to the screen with out without coloring. Most of the output functions accept an a series of string parts that are combined together. Color constants can be interlaced between these parts to color the output. 
// 
//   // vox.Println(vox.Red, "Some read text", vox.ResetColor)
//
// There are also a number of "LogLevel" type functions that easily color the output.
//
//   // vox.Alert("A warning")
//   // vox.Errorf("An error occured: %s", err.Error())
//   // vox.Debug(""A debug message")
//
// Prompting for input
//
// There are several helper functions for gathering input from the console.
//
//   // strResponse := vox.Prompt("a message", "a default value")
//   // boolResponse := vox.PromptBool("a message", true)
//   // choices := []string{"option 1", "option 2"}
//   // choiceIndex := vox.PromptChoice("A message", choices, 1)
//
//  Testing
// The output and input from for vox can be redirected to memory to make it easy to test the input and output for CLI applications. To reroute the library simply call the `Test` function.
// 
//   // vox.Test()
//
// Now data will be read/written to in memory stores instead of STDIN/STDOUT.  You can use `GetOutput` to get the current output in // the buffer. To send user input for functions like `Prompt` you can use the `SendInput` function. NOTE: SendInput must be called before any prompt function, so that the data is ready in the buffer when `Prompt` is called.

// Vox also provides an AssertOutput helper for tests that checks the current output against the passed string. It calls testing.Error if it does not match.
//
//   // func AskForFile() string {
//   //   return vox.Prompt("Enter a file", "") 
//   // }
//
//   // func TestReadConfig(t *testing.T) {
//   //   vox.Test()
//   //   
//   //   err := checkFile()
//   //   vox.AssertOutput(t, vox.Red, "No config file found.")
//   //   SetupFile("test.txt") // Builds a config file
//   //   SendInput("test.txt")
//   //   AskForFile() // Asks user for a file path  
//   //   err = checkFile()
//   //   if err != nil {
//   //     t.Errorf("Could not load config file: %s", err.Error())
//   //   }
//   //   AssertOutput(t, "Config file read")
//   // }

package vox

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Vox - The main class for Vox all functions are called from this object.
// Direct functions use an auto generated Vox object.
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

// SetOutput - Sets the output for a print actions.  By default it is Stdout.
func (v *Vox) SetOutput(w io.Writer) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.out = w
}

// SetInput - Sets the input stream for VOX. This is mainly used for testing.
func SetInput(in *os.File) { v.SetInput(in) }

// SetInput - Sets the input stream for VOX. This is mainly used for testing.
func (v *Vox) SetInput(in *os.File) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.in = in
}

// Output - Prints a string to the output stream.
func Output(s string) error { return v.Output(s) }

// Output - Prints a string to the output stream.
func (v *Vox) Output(s string) error {
	v.buf = v.buf[:0]
	v.buf = append(v.buf, s...)
	v.out.Write(v.buf)
	return nil
}

// Printf - Prints a formatted string using a template and as series of
// variables.
func Printf(format string, s ...interface{}) { v.Printf(format, s...) }

// Printf - Prints a formatted string using a template and as series of
// variables.
func (v *Vox) Printf(format string, s ...interface{}) {
	v.Output(fmt.Sprintf(format, s...))
}

// Print - Prints a number of variables.
func Print(s ...interface{}) { v.Print(s...) }

// Print - Prints a number of variables.
func (v *Vox) Print(s ...interface{}) { v.Output(fmt.Sprint(s...)) }

// Println - Prints a number of tokens ending with a new line.
func Println(s ...interface{}) { v.Println(s...) }

// Println - Prints a number of tokens ending with a new line.
func (v *Vox) Println(s ...interface{}) {
	str := fmt.Sprint(s...) + "\n"
	v.Output(str)
}

// Printlnc - Prints a number of tokens followed by a new line. This output is
// also wrapped in a color code and a reset.
func Printlnc(c Color, s ...interface{}) { v.Printlnc(c, s...) }

// Printlnc - Prints a number of tokens followed by a new line. This output is
// also wrapped in a color code and a reset.
func (v *Vox) Printlnc(c Color, s ...interface{}) {
	outStr := fmt.Sprint(s...)
	v.Println(c, outStr, ResetColor)
}

// PrintProperty - Prints a property name and value. The value will be right
// aligned.
func PrintProperty(name, value string) { v.PrintProperty(name, value) }

// PrintProperty - Prints a property name and value. The value will be right
// aligned.
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

// PrintResult - Prints a name and a result message. If an error is passed it
// will result in a failure message ex. If nil is passed as the second argument
// it will result in a success. The status code will also be right aligned and
// color coded based on the result.
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
	if err != nil {
		v.Printlnc(Red, err.Error())
	}
}

// Errorf - Print error output. Console output is colored red.
func Errorf(format string, args ...interface{}) { v.Errorf(format, args...) }

// Errorf - Print error output. Console output is colored red.
func (v *Vox) Errorf(format string, args ...interface{}) {
	v.Error(fmt.Sprintf(format, args...))
}

// Error - Print output as an error. Console output is colored red.
func Error(args ...interface{}) { v.Error(args...) }

// Error - Print output as an error. Console output is colored red.
func (v *Vox) Error(args ...interface{}) {
	v.Printlnc(Red, fmt.Sprint(args...))
}

// Infof - Print an info output. Console output is colored white.
func Infof(format string, args ...interface{}) { v.Infof(format, args...) }

// Infof - Print an info output. Console output is colored white.
func (v *Vox) Infof(format string, args ...interface{}) {
	v.Info(fmt.Sprintf(format, args...))
}

// Info - Print an info output. Console output is colored white.
func Info(args ...interface{}) { v.Info(args...) }

// Info - Print an info output. Console output is colored white.
func (v *Vox) Info(args ...interface{}) {
	v.Printlnc(White, fmt.Sprint(args...))
}

// Alertf - Print an info output. Console output is colored yellow.
func Alertf(format string, args ...interface{}) { v.Alertf(format, args...) }

// Alertf - Print an info output. Console output is colored yellow.
func (v *Vox) Alertf(format string, args ...interface{}) {
	v.Alert(fmt.Sprintf(format, args...))
}

// Alert - Print an info output. Console output is colored yellow.
func Alert(args ...interface{}) { v.Alert(args...) }

// Alert - Print an info output. Console output is colored yellow.
func (v *Vox) Alert(args ...interface{}) {
	v.Printlnc(Yellow, fmt.Sprint(args...))
}

// Debugf - Print an debug output. Debug output is not colored.
func Debugf(format string, args ...interface{}) { v.Debugf(format, args...) }

// Debugf - Print an debug output. Debug output is not colored.
func (v *Vox) Debugf(format string, args ...interface{}) {
	v.Debug(fmt.Sprintf(format, args...))
}

// Debug - Print an debug output. Debug output is not colored.
func Debug(args ...interface{}) { v.Debug(args...) }

// Debug - Print an debug output. Debug output is not colored.
func (v *Vox) Debug(args ...interface{}) {
	v.Println(fmt.Sprint(args...))
}

// Fatal - Prints an error message and then exits the application.
func Fatal(args ...interface{}) { v.Fatal(args...) }

// Fatal - Prints an error message and then exits the application.
func (v *Vox) Fatal(args ...interface{}) {
	v.Error(args...)
	os.Exit(-1)
}

// Fatalf - Prints an error message and then exits the application.
func Fatalf(format string, args ...interface{}) { v.Fatalf(format, args...) }

// Fatalf - Prints an error message and then exits the application.
func (v *Vox) Fatalf(format string, args ...interface{}) {
	v.Fatal(fmt.Sprintf(format, args...))
}
