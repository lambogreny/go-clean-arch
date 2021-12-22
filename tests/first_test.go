package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstExample(t *testing.T) {
	var a string
	var b string

	a = "Ola"
	b = "Mundo"

	assert.NotEqual(t, a, b)
	//assert.Equal(t, a, b)
}
