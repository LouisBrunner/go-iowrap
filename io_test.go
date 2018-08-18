package iowrap

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapStdout_Works(t *testing.T) {
	content := WrapStdout(func() {
		fmt.Printf("123\n")
	})
	assert.Equal(t, "123\n", string(content))
}

func TestWrapStderr_Works(t *testing.T) {
	content := WrapStderr(func() {
		fmt.Fprintf(os.Stderr, "456\n")
	})
	assert.Equal(t, "456\n", string(content))
}
