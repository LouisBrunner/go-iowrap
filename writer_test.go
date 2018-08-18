package iowrap

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter_Works(t *testing.T) {
	w, collect, err := Writer()
	if !assert.NoError(t, err) {
		return
	}

	log.SetOutput(w)
	log.Printf("hello123")
	content, err := collect()
	if assert.NoError(t, err) {
		assert.Regexp(t, " hello123\n$", string(content))
	}
}
