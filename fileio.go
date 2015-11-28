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

func WriteFile(filename string, bs []byte) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if CheckErr(err) {
		return err
	}
	defer file.Close()
	_, err = file.Write(bs)
	if CheckErr(err) {
		return err
	}
	return nil
}
