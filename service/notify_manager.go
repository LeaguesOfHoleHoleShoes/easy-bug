package service

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/g-error"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
)

const (
	MaxRecordsForPerPage = 60
)

type NotifyDB interface {
	Create(notify model.Notify) error
	LatestNotifies(proID string, count int) []model.Notify
	Notifies(proID string, page int, perPage int) (result []model.Notify, totalPages int, totalCount int)
}

func NewNotifyManager(db NotifyDB, proManager *ProjectManager) *NotifyManager {
	return &NotifyManager{db: db, proManager: proManager}
}

type NotifyManager struct {
	db NotifyDB
	proManager *ProjectManager
}

// todo 做短时间次数限制
func (m *NotifyManager) Create(proToken string, n model.Notify) error {
	if !n.NType.Valid() {
		return g_error.ErrInvalidNotifyType
	}
	p := m.proManager.GetProByToken(proToken)

	switch {
	case p.ID == "":
		return g_error.ErrInvalidProToken
	case p.Locked:
		return g_error.ErrProjectLocked
	}
	n.ProjectID = p.ID
	n.ID = util.NewObjectId()

	return m.db.Create(n)
}

func (m *NotifyManager) LatestNotifies(proToken string, count int) ([]model.Notify, error) {
	if count > MaxRecordsForPerPage {
		return nil, g_error.ErrCountMoreThanMax
	}

	p := m.proManager.GetProByToken(proToken)
	if p.ID == "" {
		return nil, g_error.ErrInvalidProToken
	}

	return m.db.LatestNotifies(p.ID, count), nil
}

func (m *NotifyManager) Notifies(proToken string, page int, perPage int) (result []model.Notify, totalPages int, totalCount int, err error) {
	if perPage > MaxRecordsForPerPage {
		err = g_error.ErrCountMoreThanMax
		return
	}

	p := m.proManager.GetProByToken(proToken)
	if p.ID == "" {
		err = g_error.ErrInvalidProToken
		return
	}

	result, totalPages, totalCount = m.db.Notifies(p.ID, page, perPage)
	return
}