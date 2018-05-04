package vox

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	Test()
	res := m.Run()
	os.Exit(res)
}

func TestOutput(t *testing.T) {
	Output("test")
	res := GetOutput()
	if res != "test" {
		t.Errorf("Result not correct: %s", res)
	}
}

func TestPrint(t *testing.T) {
	Print("OUTPUT HERE")
	res := GetOutput()
	if res != "OUTPUT HERE" {
		t.Errorf("Result not correct: %s", res)
	}
}

func TestPrintln(t *testing.T) {
	Println("OUTPUT HERE")
	res := GetOutput()
	if res != "OUTPUT HERE\n" {
		t.Errorf("Result not correct: %s", res)
	}
}

func TestPrintProperty(t *testing.T) {
	PrintProperty("Testing", "Run test")
	res := GetOutput()
	numSpaces := 60 - (len("Testing") + len("Run test"))
	expected := fmt.Sprint(Yellow, "Testing",
		strings.Repeat(" ", numSpaces), White,
		"Run test", ResetColor, "\n")

	if res != expected {
		t.Errorf("Result not correct: %s\n Expected: %s", res, expected)
	}
}

func TestSprintc(t *testing.T) {
	res := Sprintc(Red, "red")
	expected := fmt.Sprint(Red, "red", ResetColor)
	if res != expected {
		t.Errorf("Result not correct: %s\n Expected: %s", res, expected)
	}
}

func TestPrintResult(t *testing.T) {
	desc := "test"
	PrintResult(desc, nil)
	res := GetOutput()
	desc += strings.Repeat(" ", 60-len(desc))
	expected := fmt.Sprint(
		White, desc,
		Yellow, "[",
		Green, "OK",
		Yellow, "]", ResetColor, "\n")
	if res != expected {
		t.Errorf("Result not correct: %s\n Expected: %s", res, expected)
	}
}
