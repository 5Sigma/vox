package vox

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPrompt(t *testing.T) {
	ClearInput()
	SendInput("OK\n")
	result := Prompt("message", "default")
	expected := fmt.Sprintf(
		"%s%s [%s]: %s",
		Yellow,
		"message",
		"default",
		ResetColor,
	)
	if pipeline.Last() != expected {
		t.Errorf("response not correct: %s", pipeline.Last())
	}

	if result != "OK" {
		t.Errorf("Prompt response not valid: '%s'", result)
	}
}

func TestPromptBool(t *testing.T) {
	t.Run("yes response", func(t *testing.T) {
		ClearInput()
		SendInput("Y\n")
		result := PromptBool("message", false)
		expected := fmt.Sprintf(
			"%s%s [%s]: %s",
			Yellow,
			"message",
			"N",
			ResetColor,
		)

		if pipeline.Last() != expected {
			t.Errorf("response not correct: %s", pipeline.Last())
		}

		if result != true {
			t.Errorf("Prompt response not valid: '%s'", strconv.FormatBool(result))
		}
	})

	t.Run("no response", func(t *testing.T) {
		ClearInput()
		SendInput("NO\n")
		result := PromptBool("message", false)
		pipeline.Clear()

		if result != false {
			t.Errorf("Prompt response not valid: '%s'", strconv.FormatBool(result))
		}

	})

	t.Run("default response", func(t *testing.T) {
		ClearInput()
		SendInput("\n")
		result := PromptBool("message", false)
		pipeline.Clear()

		if result != false {
			t.Errorf("Prompt response not valid: '%s'", strconv.FormatBool(result))
		}
	})

}

func TestPromptChoice(t *testing.T) {
	choices := []string{
		"First",
		"Second",
		"Third",
		"Fourth",
	}
	ClearInput()
	SendInput("2")
	result := PromptChoice("message", choices, 1)
	if result != "Second" {
		t.Errorf("Prompt response not valid: '%s'", result)
	}
	ClearInput()
	SendInput("a")
	result = PromptChoice("message", choices, 0)
	if result != "First" {
		t.Errorf("Prompt response not valid: '%s'", result)
	}
	ClearInput()
	SendInput("6")
	result = PromptChoice("message", choices, 0)
	if result != "First" {
		t.Errorf("Prompt response not valid: '%s'", result)
	}
	pipeline.Clear()
	ClearInput()
}
