package vox

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

var pipeline *TestPipeline

func TestMain(m *testing.M) {
	pipeline = Test()
	res := m.Run()
	os.Exit(res)
}

func TestPrint(t *testing.T) {
	pipeline.Clear()
	Print("OUTPUT HERE")
	if pipeline.All() != "OUTPUT HERE" {
		t.Errorf("incorrect string: %s", pipeline.All())
	}
}

func TestPrintln(t *testing.T) {
	pipeline.Clear()
	Println("OUTPUT HERE")
	if pipeline.All() != "OUTPUT HERE\n" {
		t.Errorf("incorrect string: %s", pipeline.All())
	}
}

func TestPrintProperty(t *testing.T) {
	PrintProperty("Testing", "Run test")
	numSpaces := 60 - (len("Testing") + len("Run test"))
	expected := fmt.Sprint(Yellow, "Testing",
		strings.Repeat(" ", numSpaces), White,
		"Run test", ResetColor, "\n")
	if pipeline.Last() != expected {
		t.Errorf("incorrect string: \n%s%s", pipeline.Last(), expected)
	}
}

func TestSprintc(t *testing.T) {
	res := Sprintc(Red, "red")
	expected := fmt.Sprint(Red, "red", ResetColor)
	if res != expected {
		t.Errorf("incorrect string: \n%sn%s", res, expected)
	}
}

func TestPrintResult(t *testing.T) {
	t.Run("without error", func(t *testing.T) {
		pipeline.Clear()
		desc := "test"
		PrintResult(desc, nil)
		expected := fmt.Sprint(
			White, desc, strings.Repeat(" ", 60-len(desc)),
			Yellow, "[", Green, "OK", Yellow, "]",
			ResetColor, "\n",
		)
		if pipeline.Last() != expected {
			t.Errorf("incorrect string: \n%s%s", pipeline.Last(), expected)
		}
	})
	t.Run("with error", func(t *testing.T) {
		desc := "test"
		PrintResult(desc, errors.New("test error"))
		expected := fmt.Sprint(
			White,
			desc, strings.Repeat(" ", 60-len(desc)),
			Yellow, "[",
			Red, "FAIL",
			Yellow, "]", ResetColor, "\n", Red, "test error\n")
		if pipeline.Last() != expected {
			t.Errorf("incorrect string: \n%s%s", pipeline.Last(), expected)
		}
	})
}
