package iowrap

import (
	"io/ioutil"
	"os"
)

// WrapStdout will collect everything sent to the standard output
// while `functor` is executed in return it in a string
func WrapStdout(functor func()) ([]byte, error) {
	backup := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stdout = w
	defer func() {
		os.Stdout = backup
	}()
	functor()
	w.Close()
	content, err := ioutil.ReadAll(r)
	r.Close()
	if err != nil {
		return nil, err
	}
	return content, nil
}

// WrapStderr will collect everything sent to the standard error output
// while `functor` is executed in return it in a string
func WrapStderr(functor func()) ([]byte, error) {
	backup := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stderr = w
	defer func() {
		os.Stderr = backup
	}()
	functor()
	w.Close()
	content, err := ioutil.ReadAll(r)
	r.Close()
	if err != nil {
		return nil, err
	}
	return content, nil
}
