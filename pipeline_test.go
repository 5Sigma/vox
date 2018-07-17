package vox

import (
	"testing"

	"github.com/spf13/afero"
)

func TestFilePipeline(t *testing.T) {
	fs = afero.NewMemMapFs()

	pipeline := &FilePipeline{Filepath: "/var/logs/log.txt"}
	v := New()
	v.SetPipelines(pipeline)
	err := fs.MkdirAll("/var/logs", 0700)
	if err != nil {
		t.Error(err.Error())
	}
	v.Println("Something happened!!")
	v.Println("Something happened the sequel!!")
	b, err := afero.ReadFile(fs, "/var/logs/log.txt")
	if err != nil {
		t.Error(err.Error())
	}
	if string(b) != "Something happened!!\nSomething happened the sequel!!\n" {
		t.Errorf("data missmatch: %s", string(b))
	}
}
