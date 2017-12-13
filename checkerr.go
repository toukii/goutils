package goutils

import (
	"fmt"
	"log"
	// "runtime"

	cr "github.com/fatih/color"
)

var (
	infoCr = cr.New(cr.BgWhite, cr.FgRed)
	warnCr = cr.New(cr.BgYellow, cr.FgRed)
	errCr  = cr.New(cr.BgBlack, cr.FgRed)
)

func CheckErr(err error) bool {
	if nil != err {
		// funcName, file, line, ok := runtime.Caller(1)
		// if ok {
		// 	fmt.Printf("%s Line:%s Func:%s ERR:%s\n", warnCr.Sprint(file), infoCr.Sprint(line), infoCr.Sprint(runtime.FuncForPC(funcName).Name()), errCr.Sprint(err.Error()))
		// } else {
		// }
		errCr.Println(err)
		return true
	}
	return false
}

func CheckNoLogErr(err error) bool {
	if nil != err {
		return true
	}
	return false
}

// Deprecated: Use goutils.CheckErr instead.
func LogCheckErr(err error) bool {
	if nil != err {
		log.Println(err)
		return true
	}
	return false
}

func Log(v ...interface{}) {
	log.Print(v...)
}

func Print(v ...interface{}) {
	fmt.Print(v...)
}
