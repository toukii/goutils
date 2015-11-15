package goutils

import (
	"io/ioutil"
	"os"
)

func ReadFile(filename string) []byte {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if CheckErr(err) {
		return nil
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if CheckErr(err) {
		return nil
	}
	return b
}
