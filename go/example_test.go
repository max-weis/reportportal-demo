package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
}

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, nil)
}

func TestSum(t *testing.T) {
	time.Sleep(4 * time.Second)
	assert.Equal(t, 4, 2+2)
}
