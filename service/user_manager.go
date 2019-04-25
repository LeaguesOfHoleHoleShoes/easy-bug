package service

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/g-error"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"time"
)

const (
	UserTokenExpireDuration = 1 * time.Hour
)

// UserDB 用户DB
type UserDB interface {
	GetUser(id string) model.User
	GetUserByUsername(username string) model.User
	Create(u model.User) error
}

// UserManager 管理用户
type UserManager struct {
	db UserDB
}

// 暂不支持单点登录
func (m *UserManager) Login(username, password string) (result model.User, err error) {
	u := m.db.GetUserByUsername(username)
	if u.ID == "" {
		return model.User{}, g_error.ErrUsernameOrPasswordNotRight
	}
	if !u.ValidPassword(password) {
		return model.User{}, g_error.ErrUsernameOrPasswordNotRight
	}

	if u.Token, err = util.GenUserJwtToken(u.ID, UserTokenExpireDuration); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (m *UserManager) Create(u model.User) error {
	u.ID = util.NewObjectId()
	if err := u.ValidCreate(); err != nil {
		return err
	}
	u.SetEPassword()
	return m.db.Create(u)
}

// GetUserByToken 根据token获取用户
func (m *UserManager) GetUser(id string) model.User {
	return m.db.GetUser(id)
}

func (m *UserManager) ValidUserToken(token string) (u model.User, err error) {
	uID, err := util.ValidUserJwtToken(token)
	if err != nil {
		return model.User{}, err
	}

	u = m.GetUser(uID)
	if u.Username == "" {
		return model.User{}, g_error.ErrInvalidUserToken
	}

	// 这里还可以做用户付费控制等

	return u, nil
}