package test

import (
	"belajar-golang-resful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}
