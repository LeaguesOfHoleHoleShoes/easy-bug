package model

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/g-error"
)

// User 用户
type User struct {
	ID string `gorm:"unique_index" json:"id"`
	Username string `gorm:"unique_index" json:"username"`
	Password string `gorm:"-" json:"-"`
	EncryptedPassword string `json:"-"`
	// 保留字段，暂不支持 jwt token 单点登录。外边不要根据 token 来查询，应该把 token 恢复成 username 后来查
	Token string `json:"token"`

	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (u User) ValidCreate() error {
	switch {
	case u.Username == "":
		return g_error.ErrUsernameCannotEmpty
	case u.Password == "":
		return g_error.ErrPasswordCannotEmpty
	}
	return nil
}

func (u *User) SetEPassword() {
	tmpB := sha256.Sum256([]byte(u.Password))
	u.EncryptedPassword = hex.EncodeToString(tmpB[:])
}

func (u User) ValidPassword(pwd string) bool {
	tmpB := sha256.Sum256([]byte(pwd))
	if u.EncryptedPassword == hex.EncodeToString(tmpB[:]) {
		return true
	}
	return false
}