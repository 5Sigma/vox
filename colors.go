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
	Black      = Color{0}
	Red        = Color{1}
	Green      = Color{2}
	Yellow     = Color{3}
	Blue       = Color{4}
	Magenta    = Color{5}
	Cyan       = Color{6}
	White      = Color{7}
	ResetColor = Color{9}
)
