package model

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/g-error"
	"time"
)

type Project struct {
	ID string `gorm:"unique_index" json:"id"`
	Name string `json:"name"`
	Token string `gorm:"unique_index" json:"token"`
	// 标志是否可以创建 notify
	Locked bool `json:"locked"`

	UserID string `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
}

func (p Project) ValidCreate() error {
	switch {
	case p.Name == "":
		return g_error.ErrProjectNameCannotEmpty
	case p.UserID == "":
		return g_error.ErrUserIDForProjectCannotEmpty
	}
	return nil
}
