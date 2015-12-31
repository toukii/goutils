package goutils

import (
	"testing"
)

func TestWriteFile(t *testing.T) {
	WriteFile("./test/test.md", ToByte("hello,t"))
	// WriteFile("./test/test.md", ToByte("hello\ntest"))
}

func TestReWriteFile(t *testing.T) {
	ReWriteFile("./test/test.md", ToByte("hello\ntest"))
	ReWriteFile("./test/test.md", nil)
}
