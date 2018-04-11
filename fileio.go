package goutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFile(filename string, logg ...bool) []byte {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if logg == nil && CheckErr(err) || logg != nil && logg[0] && err != nil {
		return nil
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if logg == nil && CheckErr(err) || logg != nil && logg[0] && err != nil {
		return nil
	}
	return b
}

func DeleteFile(filename string) error {
	info, err := os.Stat(filename)
	if info != nil || err == nil {
		err = os.Remove(filename)
		if CheckNoLogErr(err) {
			return err
		}
	}
	return err
}

func FileIsExists(filename string) bool {
	info, err := os.Stat(filename)
	if info != nil || err == nil {
		return true
	}
	return false
}

func Mkdir(dirname string) error {
	dinfo, err := os.Stat(dirname)
	if err != nil || dinfo == nil {
		if err := (os.MkdirAll(dirname, 0777)); err != nil {
			return err
		}
	}
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
	fmt.Printf("data writing ==> %s\n", filename)
	return nil
}

func SafeWriteFile(filename string, bs []byte) error {
	dir := filepath.Dir(filename)
	dinfo, err := os.Stat(dir)
	if err != nil || dinfo == nil {
		if err := (os.MkdirAll(dir, 0777)); err != nil {
			return err
		}
	}
	if FileIsExists(filename) {
		return fmt.Errorf("%s already exists.", filename)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if CheckErr(err) {
		return err
	}
	defer file.Close()
	_, err = file.Write(bs)
	if CheckErr(err) {
		return err
	}
	fmt.Printf("data writing ==> %s\n", filename)
	return nil
}

func ReWriteFile(filename string, bs []byte) error {
	return ioutil.WriteFile(filename, bs, os.ModeAppend)
}
