package mysql

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

func NewProjectDB(db *gorm.DB) *ProjectDB {
	return &ProjectDB{db: db}
}

type ProjectDB struct {
	db *gorm.DB
}

func (db *ProjectDB) GetProjects(uID string, page int, perPage int) (result []model.Project, totalPages int, totalCount int) {
	fDB := db.db.Where("user_id=?", uID)
	totalPages, totalCount = util.GetDataByPageAndPerPage(fDB, page, perPage, model.Project{}, &result)
	return
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
	p.CreatedAt = time.Now()
	return db.db.Model(model.Project{}).Create(p).Error
}

