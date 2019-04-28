package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
)

func TestUserDB_GetUser(t *testing.T) {
	defer gDB.ClearAllData()

	uDB := NewUserDB(gDB.GetDB())
	assert.NoError(t, uDB.Create(model.User{
		ID: "u1",
		Username: "u122",
		EncryptedPassword: "123",
		// 保留字段，暂不支持 jwt token 单点登录。外边不要根据 token 来查询，应该把 token 恢复成 username 后来查
		Token: "321",

		Name: "qqa",
		Phone: "111",

		CreatedAt: time.Now(),
	}))

	u := uDB.GetUser("u1")
	assert.NotEmpty(t, u.ID)
	u = uDB.GetUserByUsername("u122")
	assert.NotEmpty(t, u.ID)
}
