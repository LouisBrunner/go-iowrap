package iowrap

import (
	"io/ioutil"
	"os"
)

// WrapStdout will collect everything sent to the standard output
// while `functor` is executed in return it in a string
func WrapStdout(functor func()) []byte {
	backup := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = backup
	}()
	functor()
	w.Close()
	content, _ := ioutil.ReadAll(r)
	return content
}

// WrapStderr will collect everything sent to the standard error output
// while `functor` is executed in return it in a string
func WrapStderr(functor func()) []byte {
	backup := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w
	defer func() {
		os.Stderr = backup
	}()
	functor()
	w.Close()
	content, _ := ioutil.ReadAll(r)
	return content
}
