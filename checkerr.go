package goutils

import (
	"log"
)

func CheckErr(err error) bool {
	if nil != err {
		log.Println(err)
		return true
	}
	return false
}
