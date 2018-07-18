package vox

import (
	"fmt"
)

// Color - A structure which represents a single color. This structure should
// not need to be used directly. Variables for each color are exported in the
// package.
type Color struct {
	index int
}

func (c Color) String() string {
	return fmt.Sprintf("\u001b[%dm", 30+c.index)
}

var (
	// Back - Black terminal color constant. This can be used as a string inside any of the output functions
	Black = Color{0}

	// Red - Red terminal color constant. This can be used as a string inside any of the output functions
	Red = Color{1}

	// Green - Green terminal color constant. This can be used as a string inside any of the output functions
	Green = Color{2}

	// Yellow - Yellow terminal color constant. This can be used as a string inside any of the output functions
	Yellow = Color{3}

	// Blue - Blue terminal color constant. This can be used as a string inside any of the output functions
	Blue = Color{4}

	// Magenta - Magenta terminal color constant. This can be used as a string inside any of the output functions
	Magenta = Color{5}

	// Cyan - Cyan terminal color constant. This can be used as a string inside any of the output functions
	Cyan = Color{6}

	// White - White terminal color constant. This can be used as a string inside any of the output functions
	White = Color{7}

	// ResetColor - Resets the terminal back to its default color. This can be used as a string inside any of the output functions
	ResetColor = Color{9}
)
