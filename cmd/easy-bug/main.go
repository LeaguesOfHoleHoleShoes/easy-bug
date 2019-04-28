package main

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	mysql2 "github.com/LeaguesOfHoleHoleShoes/easy-bug/db/mysql"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/common"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/gin-handler"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/rpcx-handler"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/service"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
)

/*

环境变量

docker_env: 0 非docker环境 1 docker中非生产环境 2 docker中生产环境 默认 0
db_env: dev test preprod prod
run_env: dev test preprod prod

gin_port: 默认3001
rpcx_port: 默认3000

create_user_token: 创建用户的暗号

*/
func main() {
	db := mysql.MakeDB(mysql2.GetDBConfig())
	nDB := mysql2.NewNotifyDB(db.GetDB())
	pDB := mysql2.NewProjectDB(db.GetDB())
	uDB := mysql2.NewUserDB(db.GetDB())
	pManager := service.NewProjectManager(pDB)
	nManager := service.NewNotifyManager(nDB, pManager)
	uManager := service.NewUserManager(uDB)
	env := common.NewEnv(uManager, pManager, nManager)
	handler := common.NewHandler(*env)

	go func() {
		ginS := gin_handler.NewServer(handler)
		if err := ginS.Run(":" + util.GetEnv("gin_port", "3000")); err != nil {
			panic(err)
		}
	}()

	rpcxS := rpcx_handler.NewServer(handler)
	if err := rpcxS.Serve("tcp", ":" + util.GetEnv("rpcx_port", "3001")); err != nil {
		panic(err)
	}
}
