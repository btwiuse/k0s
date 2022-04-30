package errors

import (
	"fmt"
	"runtime"
	"strings"
)

func getPkg() string {
	var idx int
	var pkg string
	var pc uintptr

	pc, _, _, _ = runtime.Caller(2)
	pkg = runtime.FuncForPC(pc).Name()

	if (pkg == "") || strings.HasPrefix(pkg, "main.") {
		return ""
	}

	idx = strings.LastIndex(pkg, "/")
	pkg = pkg[idx+1:]

	idx = strings.Index(pkg, ".")
	if idx < 0 {
		return pkg
	}

	return pkg[:idx] + ": "
}

// New will return a new error with a prefixed package name.
func New(str string) error {
	return fmt.Errorf(getPkg() + str)
}

// Newf will return a new error from format string with a prefixed
// package name.
func Newf(str string, params ...interface{}) error {
	return fmt.Errorf(getPkg()+str, params...)
}
