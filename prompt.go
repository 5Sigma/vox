package vox

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Prompt - Gets input from the input stream. By default Stdin. If an empty
// string is sent the default value will be returned.
func Prompt(name, defaultVal string) string { return v.Prompt(name, defaultVal) }

// Prompt - Gets input from the input stream. By default Stdin. If an empty
// string is sent the default value will be returned.
func (v *Vox) Prompt(name, defaultValue string) string {
	reader := bufio.NewReader(v.in)
	if defaultValue != "" {
		Printf("%s%s [%s]: %s", Yellow, name, defaultValue, ResetColor)
	} else {
		Printf("%s%s : %s", Yellow, name, ResetColor)
	}
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" && defaultValue != "" {
		return defaultValue
	}
	return input
}

// PromptBool - Prompts the user for a boolean response.
func PromptBool(message string, defaultVal bool) bool {
	return v.PromptBool(message, defaultVal)
}

// PromptBool - Prompts the user for a boolean response.
func (v *Vox) PromptBool(message string, defaultVal bool) bool {

	defaultValStr := "Y"
	if !defaultVal {
		defaultValStr = "N"
	}

	Printf("%s%s [%s]: %s", Yellow, message, defaultValStr, ResetColor)

	reader := bufio.NewReader(v.in)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))
	retValue := defaultVal

	if input == "y" || input == "yes" {
		retValue = true
	}

	if input == "n" || input == "no" {
		retValue = false
	}

	return retValue
}

// PromptChoice - Prompts for a choice of a series of options from the user.
func PromptChoice(msg string, choices []string, defaultIdx int) string {
	return v.PromptChoice(msg, choices, defaultIdx)
}

// PromptChoice - Prompts for a choice of a series of options from the user.
func (v *Vox) PromptChoice(msg string, choices []string, defIdx int) string {
	output := []string{"Choose an option:"}
	for idx, c := range choices {
		output = append(output, fmt.Sprintf("%d. %s", idx+1, c))
	}
	Printf(strings.Join(output, "\n"))
	Print(Yellow, "[", choices[defIdx], "] ", ResetColor)
	reader := bufio.NewReader(v.in)
	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(input)
	println(choice)
	if err != nil {
		return choices[defIdx]
	}
	if choice < len(choices) {
		return choices[choice-1]
	}
	return choices[defIdx]
}
