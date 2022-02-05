package main //nolint:testpackage // It is the main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, 6, Calculate(3))
}
