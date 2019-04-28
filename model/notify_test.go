package model

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNotifyType(t *testing.T) {
	type A struct {
		X NotifyType `json:"x"`
	}
	var a A
	assert.NoError(t, util.ParseJsonFromBytes(util.StringifyJsonToBytes(map[string]string{
		"x": "123",
	}), &a))
	assert.Equal(t, NotifyType("123"), a.X)
}
