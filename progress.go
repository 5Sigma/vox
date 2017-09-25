package vox

import (
	"fmt"
	"github.com/gosuri/uilive"
	"strings"
	"time"
)

// ProgressBar - A structure that controls displaying a progress bar in the
// console.
type progress struct {
	Writer    *uilive.Writer
	Max       int
	Current   int
	StartTime time.Time
}

// StartProgress - Start outputing a progressbar.
func StartProgress(current, max int) { v.StartProgress(current, max) }
func (v *Vox) StartProgress(current, max int) {
	v.progress = &progress{
		Writer:    uilive.New(),
		Max:       max,
		Current:   current,
		StartTime: time.Now(),
	}
	v.progress.Writer.Out = v.out
	v.progress.Writer.Start()
}

// IncProgress - Increment the current progress value by name. If the new
// Current value is equal to the Max value StopProgress will be called
// automatically.
func IncProgress() { v.IncProgress() }
func (v *Vox) IncProgress() {
	v.progress.Current++
	if v.progress.Current == v.progress.Max {
		v.StopProgress()
	}
	v.writeProgress()
}

// SetProgress - Sets the current progress value.
func SetProgress(current int) { v.SetProgress(current) }
func (v *Vox) SetProgress(current int) {
	v.progress.Current = current
	if v.progress.Current >= v.progress.Max {
		v.StopProgress()
	}
	v.writeProgress()
}

func (v *Vox) writeProgress() {
	elapsed := time.Since(v.progress.StartTime)
	perc := (float64(v.progress.Current) / float64(v.progress.Max)) * float64(10)
	barStr := strings.Repeat("-", 10)
	barStr = strings.Replace(barStr, "-", "=", int(perc))
	line := fmt.Sprintf("[%d/%d] %s %s", v.progress.Current, v.progress.Max,
		barStr, elapsed)
	fmt.Fprintln(v.progress.Writer, line)
}

// StopProgress - Stops outputing a progress bar and closes associated writers.
// This is called automatically if the Current value equals, or exceeds, the
// maximum value.
func StopProgress() { v.StopProgress() }
func (v *Vox) StopProgress() {
	v.progress.Writer.Stop()
}
