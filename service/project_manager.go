package service

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/hashicorp/golang-lru"
)

type ProjectDB interface {
	GetProByToken(token string) model.Project
	Create(p model.Project) error
	GetProjects(uID string, page int, perPage int) (result []model.Project, totalPages int, totalCount int)
}

func NewProjectManager(db ProjectDB) *ProjectManager {
	c, err := lru.New(800)
	if err != nil {
		panic(err)
	}
	return &ProjectManager{
		db: db,
		projectCache: c,
	}
}

type ProjectManager struct {
	db ProjectDB
	projectCache *lru.Cache
}

func (m *ProjectManager) Create(p model.Project) (result model.Project, err error) {
	p.ID = util.NewObjectId()
	if p.Token, err = util.GenRandomToken(); err != nil {
		return model.Project{}, err
	}

	// 验证是否合法
	if err = p.ValidCreate(); err != nil {
		return
	}
	return p, m.db.Create(p)
}

func (m *ProjectManager) GetProjects(uID string, page int, perPage int) (result []model.Project, totalPages int, totalCount int, err error) {
	result, totalPages, totalCount = m.db.GetProjects(uID, page, perPage)
	return
}

// 外边要检查用户的 token ?
func (m *ProjectManager) GetProByToken(token string) model.Project {
	if p, ok := m.projectCache.Get(token); ok {
		return p.(model.Project)
	}

	p := m.db.GetProByToken(token)
	if p.ID != "" {
		m.projectCache.Add(token, p)
	}

	return p
}
