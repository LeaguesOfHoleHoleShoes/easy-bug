package mysql

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

type UserDB struct {
	db *gorm.DB
}

func (db *UserDB) GetUser(id string) (result model.User) {
	if err := db.db.Model(model.User{}).First(&result, "id=?", id).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.QyLogger.Warn("can't get user by id", zap.Error(err))
		}
	}
	return
}

func (db *UserDB) GetUserByUsername(username string) (result model.User) {
	if err := db.db.Model(model.User{}).First(&result, "username=?", username).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.QyLogger.Warn("can't get user by username", zap.Error(err))
		}
	}
	return
}

func (db *UserDB) Create(u model.User) error {
	u.CreatedAt = time.Now()
	return db.db.Model(model.User{}).Create(u).Error
}
