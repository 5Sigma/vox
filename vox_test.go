package vox

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var out = new(bytes.Buffer)

func TestMain(m *testing.M) {
	SetOutput(out)
	res := m.Run()
	os.Exit(res)
}

func TestOutput(t *testing.T) {
	Output("test")
	res, _ := out.ReadString('\n')
	assert.Equal(t, "test", res)
}

func TestPrint(t *testing.T) {
	Print("OUTPUT HERE")
	res, _ := out.ReadString('\n')
	assert.Equal(t, "OUTPUT HERE", res)
}

func TestPrintln(t *testing.T) {
	Println("OUTPUT HERE")
	res, _ := out.ReadString('\n')
	assert.Equal(t, "OUTPUT HERE", res)
}

func TestPrintProperty(t *testing.T) {
	PrintProperty("Testing", "Run test")
	res, _ := out.ReadString('\n')
	numSpaces := 60 - (len("Testing") + len("Run test"))
	expected := fmt.Sprint(Yellow, "Testing",
		strings.Repeat(" ", numSpaces), White,
		"Run test", ResetColor)
	assert.Equal(t, expected, res)
}

func TestSprintc(t *testing.T) {
	res := Sprintc(Red, "red")
	expected := fmt.Sprint(Red, "red", ResetColor)
	assert.Equal(t, expected, res)
}

func TestPrintResult(t *testing.T) {
	desc := "test"
	PrintResult(desc, nil)
	res, _ := out.ReadString('\n')
	desc += strings.Repeat(" ", 60-len(desc))
	expected := fmt.Sprint(
		White, desc,
		Yellow, "[",
		Green, "OK",
		Yellow, "]", ResetColor)
	assert.Equal(t, expected, res)
}
