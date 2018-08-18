package iowrap

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapStdout_Works(t *testing.T) {
	content, err := WrapStdout(func() {
		fmt.Printf("123\n")
	})
	if assert.NoError(t, err) {
		assert.Equal(t, "123\n", string(content))
	}
}

func TestWrapStderr_Works(t *testing.T) {
	content, err := WrapStderr(func() {
		fmt.Fprintf(os.Stderr, "456\n")
	})
	if assert.NoError(t, err) {
		assert.Equal(t, "456\n", string(content))
	}
}
