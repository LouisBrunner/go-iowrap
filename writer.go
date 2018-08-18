package iowrap

import (
	"io"
	"io/ioutil"
	"os"
)

// Writer returns a writer that will collect all output and which can then be
// accessed by calling the second return (which will also close the writer)
func Writer() (io.Writer, func() ([]byte, error), error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return w, func() ([]byte, error) {
		w.Close()
		content, err := ioutil.ReadAll(r)
		r.Close()
		if err != nil {
			return nil, err
		}
		return content, nil
	}, nil
}
