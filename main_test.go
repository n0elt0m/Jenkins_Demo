package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomNumber(t *testing.T) {
	n := RandomNumber()
	assert.GreaterOrEqual(t, n, 0)
	assert.LessOrEqual(t, n, 1000)
}
