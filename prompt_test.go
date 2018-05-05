package vox

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPrompt(t *testing.T) {
	SendInput("OK\n")
	result := Prompt("message", "default")
	AssertOutput(t,
		fmt.Sprintf(
			"%s%s [%s]: %s",
			Yellow,
			"message",
			"default",
			ResetColor,
		),
	)

	if result != "OK" {
		t.Errorf("Prompt response not valid: '%s'", result)
	}
}

func TestPromptBool(t *testing.T) {
	ClearInput()
	SendInput("Y\n")
	result := PromptBool("message", false)
	AssertOutput(t,
		fmt.Sprintf(
			"%s%s [%s]: %s",
			Yellow,
			"message",
			"N",
			ResetColor,
		),
	)

	if result != true {
		t.Errorf("Prompt response not valid: '%s'", strconv.FormatBool(result))
	}

	ClearInput()
	SendInput("NO\n")
	result = PromptBool("message", false)
	ClearOutput()

	if result != false {
		t.Errorf("Prompt response not valid: '%s'", strconv.FormatBool(result))
	}

	ClearInput()
	SendInput("\n")
	result = PromptBool("message", false)
	ClearOutput()

	if result != false {
		t.Errorf("Prompt response not valid: '%s'", strconv.FormatBool(result))
	}
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
	ClearOutput()
	ClearInput()
}
