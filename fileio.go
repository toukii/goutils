package goutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
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

func DeleteFile(filename string) error {
	info, err := os.Stat(filename)
	if info != nil || err == nil {
		err = os.Remove(filename)
		if CheckErr(err) {
			return err
		}
	}
	CheckErr(err)
	return nil
}

func WriteFile(filename string, bs []byte) error {
	dir := filepath.Dir(filename)
	dinfo, err := os.Stat(dir)
	if err != nil || dinfo == nil {
		if err := (os.MkdirAll(dir, 0777)); err != nil {
			return err
		}
	}
	DeleteFile(filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
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

func ReWriteFile(filename string, bs []byte) error {
	return ioutil.WriteFile(filename, bs, os.ModeAppend)
}
