package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	p1, err := parser("TX13ABCDEFLMSSSSS")
	assert.Equal(t, len([]rune(p1.Value)), p1.Length, err)
	p2, err := parser("TXY6ABCDE")
	assert.Equal(t, len([]rune(p2.Value)), p2.Length, err)
	p3, err := parser("NK04000GA")
	assert.Equal(t, len([]rune(p3.Value)), p3.Length, err)
	p4, err := parser("NN04000GADFS")
	assert.Equal(t, len([]rune(p4.Value)), p4.Length, err)
}
