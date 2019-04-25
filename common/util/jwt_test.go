package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenProjectJwtToken(t *testing.T) {
	token, err := GenUserJwtToken("123", time.Second)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidUserJwtToken(t *testing.T) {
	// expire 记录的是秒
	token, err := GenUserJwtToken("123", 2 * time.Second)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// 测试没超时
	uID, err := ValidUserJwtToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "123", uID)

	// 测试超时
	time.Sleep(3 * time.Second)
	uID, err = ValidUserJwtToken(token)
	assert.Error(t, err)
	assert.Equal(t, "", uID)
}
