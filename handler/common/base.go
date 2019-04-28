package common

import "github.com/LeaguesOfHoleHoleShoes/easy-bug/service"

func NewEnv(uManager *service.UserManager, pManager *service.ProjectManager, nManager *service.NotifyManager) *Env {
	return &Env{uManager: uManager, pManager: pManager, nManager: nManager}
}

type Env struct {
	uManager *service.UserManager
	pManager *service.ProjectManager
	nManager *service.NotifyManager
}
