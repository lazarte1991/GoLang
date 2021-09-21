package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	p1 := parser("TX13ABCDEFLMSSSSS")
	assert.Equal(t, len([]rune(p1.Value)), p1.Length, "los valores no coinciden")
	p2 := parser("TX06ABCDE")
	assert.Equal(t, len([]rune(p2.Value)), p2.Length, "los valores no coinciden")
	p3 := parser("NN04000GA")
	assert.Equal(t, len([]rune(p3.Value)), p3.Length, "los valores no coinciden")
}
