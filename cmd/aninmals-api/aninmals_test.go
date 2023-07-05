package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAninmals(t *testing.T) {
	result := testFunction()
	assert.Equal(t, 666, result)
}
