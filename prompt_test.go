package vox

import (
	"fmt"
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
