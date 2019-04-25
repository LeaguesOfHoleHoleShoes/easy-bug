package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenRandomToken()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
