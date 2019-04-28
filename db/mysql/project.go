package mysql

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type ProjectDB struct {
	db *gorm.DB
}

func (db *ProjectDB) GetProByToken(token string) (result model.Project) {
	if err := db.db.Model(model.Project{}).First(&result, "token=?", token).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.QyLogger.Warn("get pro by token failed", zap.Error(err))
		}
	}
	return
}

func (db *ProjectDB) Create(p model.Project) error {
	return db.db.Model(model.Project{}).Create(p).Error
}

