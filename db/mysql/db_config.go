package mysql

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/g-env"
	"github.com/dipperin/go-ms-toolkit/db-config"
	_ "github.com/go-sql-driver/mysql"
)

func GetDBConfig() *db_config.DbConfig {
	conf := db_config.NewDbConfig()
	dbName := "easy_bug"

	runEnv := g_env.GetDBEnv()
	switch runEnv {
	case "test", "preprod", "prod":
		dbName += "_" + runEnv
	default:
		dbName += "_dev"
	}

	conf.DbName = dbName
	return conf
}
