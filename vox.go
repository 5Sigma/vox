package vox

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// Vox - The main class for Vox all functions are called from this object.
// Direct functions use an auto generated Vox object.
type Vox struct {
	mu        sync.Mutex
	buf       []byte
	in        *os.File
	progress  *progress
	pipelines []Pipeline
}

var v *Vox

func init() {
	v = New()
}

// New - creates a new Vox instance. This can be used as an alternative to the
// singletone instance. If multiple Vox instances are needed.
func New() *Vox {
	v := &Vox{}
	v.SetPipelines(&ConsolePipeline{})
	return v
}

// Write writes data into the log
func (v *Vox) Write(p []byte) (n int, err error) {
	for _, pl := range v.pipelines {
		_, err := pl.Write(p)
		if err != nil {
			return 0, err
		}
	}
	return len(p), nil
}

// Sprintc - Creates a string in a given color. The color code prepends the
// string and a reset code is appended to it.
func Sprintc(c Color, args ...interface{}) string {
	str := fmt.Sprint(args...)
	return fmt.Sprint(c, str, ResetColor)
}

// SetPipelines replaces all pipelines with the passed pipeline
func (v *Vox) SetPipelines(p Pipeline) {
	if err := p.Initialize(); err == nil {
		v.pipelines = []Pipeline{p}
	} else {
		v.pipelines = []Pipeline{}
	}
}

// SetPipelines replaces all pipelines with the passed pipeline
func SetPipelines(p Pipeline) { v.SetPipelines(p) }

// AddPipeline adds a new pipeline to the logger
func (v *Vox) AddPipeline(p Pipeline) {
	if err := p.Initialize(); err == nil {
		v.pipelines = append(v.pipelines, p)
	}
}

// AddPipeline adds a new pipeline to the logger
func AddPipeline(p Pipeline) { v.AddPipeline(p) }

// SetInput - Sets the input stream for VOX. This is mainly used for testing.
func SetInput(in *os.File) { v.SetInput(in) }

// SetInput - Sets the input stream for VOX. This is mainly used for testing.
func (v *Vox) SetInput(in *os.File) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.in = in
}

func (v *Vox) output(s string) error {
	for _, pl := range v.pipelines {
		if !pl.Config().Plain {
			v.buf = v.buf[:0]
			v.buf = append(v.buf, s...)

			pl.Write(v.buf)
			_, err := pl.Write(v.buf)
			if err != nil {
				println(err.Error())
			}
		}
	}
	return nil
}

func (v *Vox) outputPlain(s string) error {
	for _, pl := range v.pipelines {
		if pl.Config().Plain {
			v.buf = v.buf[:0]
			v.buf = append(v.buf, s...)

			_, err := pl.Write(v.buf)
			if err != nil {
				println(err.Error())
			}
		}
	}
	return nil
}

// Printf - Prints a formatted string using a template and as series of
// variables.
func Printf(format string, s ...interface{}) { v.Printf(format, s...) }

// Printf - Prints a formatted string using a template and as series of
// variables.
func (v *Vox) Printf(format string, s ...interface{}) {
	v.output(fmt.Sprintf(format, s...))
}

// Print - Prints a number of variables.
func Print(s ...interface{}) { v.Print(s...) }

// Print - Prints a number of variables.
func (v *Vox) Print(s ...interface{}) {
	v.output(fmt.Sprint(s...))
	v.outputPlain(fmt.Sprint(s...))
}

// Println - Prints a number of tokens ending with a new line.
func Println(s ...interface{}) { v.Println(s...) }

// Println - Prints a number of tokens ending with a new line.
func (v *Vox) Println(s ...interface{}) {
	str := fmt.Sprint(s...) + "\n"
	v.Print(str)
}

// Printlnc - Prints a number of tokens followed by a new line. This output is
// also wrapped in a color code and a reset.
func Printlnc(c Color, s ...interface{}) { v.Printlnc(c, s...) }

// Printlnc - Prints a number of tokens followed by a new line. This output is
// also wrapped in a color code and a reset.
func (v *Vox) Printlnc(c Color, s ...interface{}) {
	outStr := fmt.Sprint(s...)
	v.output(fmt.Sprint(c, outStr, ResetColor, "\n"))
	v.outputPlain(outStr + "\n")
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
	var (
		out      string
		outPlain string
	)
	resultColor := Red
	resultText := "FAIL"
	if err == nil {
		resultColor = Green
		resultText = "OK"
	}
	desc += strings.Repeat(" ", 60-len(desc))
	out += fmt.Sprint(
		White, desc,
		Yellow, "[", resultColor, resultText, Yellow, "]",
		ResetColor,
		"\n",
	)
	if err != nil {
		out += fmt.Sprint(Red, err.Error(), "\n")
	}
	v.output(out)

	outPlain += fmt.Sprintf("%s [%s]\n", desc, resultText)
	if err != nil {
		outPlain += err.Error() + "\n"
	}

	v.outputPlain(outPlain)
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
