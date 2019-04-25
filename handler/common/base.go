package common

import "github.com/LeaguesOfHoleHoleShoes/easy-bug/service"

type Env struct {
	uManager *service.UserManager
	pManager *service.ProjectManager
	nManager *service.NotifyManager
}
