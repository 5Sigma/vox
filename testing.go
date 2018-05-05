package vox

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// Test - Sets up vox to print and read from in memory locations for testing.
func Test() error { return v.Test() }

// Test - Sets up vox to print and read from in memory locations for testing.
func (v *Vox) Test() error {
	in, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	out := new(bytes.Buffer)
	if err != nil {
		return err
	}
	v.SetOutput(out)
	v.SetInput(in)
	return nil
}

// GetOutput - Returns the output in the buffer used for testing.
func GetOutput() string { return v.GetOutput() }

// GetOutput - Returns the output in the buffer used for testing.
func (v *Vox) GetOutput() string {
	res, _ := v.out.(*bytes.Buffer).ReadString('\n')
	return res
}

// SendInput - Writes data into the input stream. Used for testing.
func SendInput(str string) error { return v.SendInput(str) }

// ClearInput - Clears the input buffer. Useful during testing.
func ClearInput() {
	in, _ := ioutil.TempFile("", "")
	v.SetInput(in)
}

// ClearOutput - Clears the ouput buffer. Useful during testing.
func ClearOutput() {
	out := new(bytes.Buffer)
	v.SetOutput(out)
}

// SendInput - Writes data into the input stream. Used for testing.
func (v *Vox) SendInput(str string) error {
	_, err := io.WriteString(v.in, str)
	if err != nil {
		return err
	}
	v.in.Seek(0, os.SEEK_SET)
	return nil
}

// AssertOutput - Checks the output for a given string and throws a testing
// error if it does not match.
func AssertOutput(t *testing.T, args ...interface{}) {
	res := GetOutput()
	str := fmt.Sprint(args...)
	if res != str {
		t.Errorf("Expected: '%s'\nRecieved: '%s'", str, res)
	}
}
