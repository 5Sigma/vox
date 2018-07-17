package vox

import (
	"io"
	"os"
)

// PipelineConfig is the configuration for a pipeline
// The pipeline config should be returned from the Config function for any
// pipeline structure.
type PipelineConfig struct {
	// Plain - If set to true all color information and some formatting will be
	// stripped from the output. This is normally used for outputs directed towards
	// files.
	Plain bool
}

// Pipeline represents a specific log pipeline
type Pipeline interface {
	Config() *PipelineConfig
	Write([]byte) (int, error)
	Initialize() error
}

// ConsolePipeline a log pipeline that outputs directly to STDERR
type ConsolePipeline struct{}

// Config returns the pipeline configuration
func (c *ConsolePipeline) Config() *PipelineConfig {
	return &PipelineConfig{
		Plain: false,
	}
}

// Write sends data to the output Stdout
func (c *ConsolePipeline) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

// Initialize has no logic for a ConsolePipeline
func (c *ConsolePipeline) Initialize() error {
	return nil
}

// FilePipeline sends output into a local file
type FilePipeline struct {
	Filepath string
	file     *os.File
}

// Config returns the pipline configuration
func (f *FilePipeline) Config() *PipelineConfig {
	return &PipelineConfig{
		Plain: true,
	}
}

// Write sends the data to the local filepath
func (f *FilePipeline) Write(b []byte) (int, error) {
	return f.file.Write(b)
}

// Initialize opens the local file for reading
func (f *FilePipeline) Initialize() error {
	var err error
	f.file, err = os.Open(f.Filepath)
	return err
}

// TestPipeline a pipeline that can be used in tests
type TestPipeline struct {
	LogLines []string
	Plain    bool
}

// Config returns the pipline configuration
func (t *TestPipeline) Config() *PipelineConfig {
	return &PipelineConfig{
		Plain: t.Plain,
	}
}

func (t *TestPipeline) Write(b []byte) (int, error) {
	t.LogLines = append(t.LogLines, string(b))
	return len(b), nil
}

// Initialize sets up the testing pipeline
func (t *TestPipeline) Initialize() error {
	t.LogLines = []string{}
	return nil
}

// Last returns the last section of data sent
func (t *TestPipeline) Last() string {
	if len(t.LogLines) == 0 {
		return ""
	}
	return t.LogLines[len(t.LogLines)-1]
}

// Clear removes all items in the pipelines buffer
func (t *TestPipeline) Clear() {
	t.LogLines = []string{}
}

// WriterPipeline implements a generic pipeline powered by an io.Writer stream
type WriterPipeline struct {
	Writer io.Writer
	Plain  bool
}

// Config returns a configuration for the pipeline. Plain is specified on the
// pipeline itself and patched into the configuration.
func (w *WriterPipeline) Config() *PipelineConfig {
	return &PipelineConfig{Plain: w.Plain}
}

// Write sends data into the specified writer.
func (w *WriterPipeline) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Initialize has no logic in a WriterPipeline
func (w *WriterPipeline) Initialize() error {
	return nil
}
