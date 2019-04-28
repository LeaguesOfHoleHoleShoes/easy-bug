package mysql

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type NotifyDB struct {
	db *gorm.DB
}

func (db *NotifyDB) Create(notify model.Notify) error {
	return db.db.Model(model.Notify{}).Create(notify).Error
}

func (db *NotifyDB) LatestNotifies(proID string, count int) (result []model.Notify) {
	if err := db.db.Model(model.Notify{}).Order("created_at desc").Limit(count).Find(&result).Error; err != nil {
		log.QyLogger.Warn("find latest notifies failed", zap.Error(err))
	}
	return
}

func (db *NotifyDB) Notifies(proID string, page int, perPage int) (result []model.Notify, totalPages int, totalCount int) {
	fDB := db.db.Where("project_id=?", proID)
	totalPages, totalCount = util.GetDataByPageAndPerPage(fDB, page, perPage, model.Notify{}, &result)
	return
}



