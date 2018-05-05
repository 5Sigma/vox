package vox

import (
	"errors"
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
	desc += strings.Repeat(" ", 60-len(desc))
	PrintResult(desc, nil)
	AssertOutput(t,
		White, desc,
		Yellow, "[", Green, "OK", Yellow, "]",
		ResetColor, "\n",
	)
	ClearOutput()
	PrintResult(desc, errors.New("test error"))
	AssertOutput(t,
		White,
		desc,
		Yellow, "[",
		Red, "FAIL",
		Yellow, "]", ResetColor, "\n",
	)
	ClearOutput()
}

func TestErrorf(t *testing.T) {
	ClearOutput()
	Errorf("test error")
	AssertOutput(t, Red, "test error", ResetColor, "\n")
	ClearOutput()

}

func TestInfof(t *testing.T) {
	ClearOutput()
	Infof("test info")
	AssertOutput(t, White, "test info", ResetColor, "\n")
	ClearOutput()
}

func TestAlertf(t *testing.T) {
	ClearOutput()
	Alertf("test alert")
	AssertOutput(t, Yellow, "test alert", ResetColor, "\n")
	ClearOutput()
}

func TestDebugf(t *testing.T) {
	ClearOutput()
	Debugf("test debug")
	AssertOutput(t, "test debug", "\n")
	ClearOutput()
}

func TestError(t *testing.T) {
	ClearOutput()
	Error("test error")
	AssertOutput(t, Red, "test error", ResetColor, "\n")
	ClearOutput()

}

func TestInfo(t *testing.T) {
	ClearOutput()
	Info("test info")
	AssertOutput(t, White, "test info", ResetColor, "\n")
	ClearOutput()
}

func TestAlert(t *testing.T) {
	ClearOutput()
	Alert("test alert")
	AssertOutput(t, Yellow, "test alert", ResetColor, "\n")
	ClearOutput()
}

func TestDebug(t *testing.T) {
	ClearOutput()
	Debug("test debug")
	AssertOutput(t, "test debug", "\n")
	ClearOutput()
}
