package vox

import (
	"io"
	"io/ioutil"
	"os"
)

// Test - Sets up vox to print and read from in memory locations for testing.
func Test() *TestPipeline { return v.Test() }

// Test - Sets up vox to print and read from in memory locations for testing.
func (v *Vox) Test() *TestPipeline {
	pl := &TestPipeline{}
	v.SetPipelines(pl)
	return pl
}

// SendInput - Writes data into the input stream. Used for testing.
func SendInput(str string) error { return v.SendInput(str) }

// ClearInput - Clears the input buffer. Useful during testing.
func ClearInput() {
	in, _ := ioutil.TempFile("", "")
	v.SetInput(in)
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
