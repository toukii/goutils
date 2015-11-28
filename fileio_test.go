package goutils

import (
	"testing"
)

func TestWriteFile(t *testing.T) {
	WriteFile("./test/test.md", ToByte("hello,test"))
}
