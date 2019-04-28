package main

import (
	mysql2 "github.com/LeaguesOfHoleHoleShoes/easy-bug/db/mysql"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "Easy Bug Init Mysql DB"
	app.Commands = []cli.Command{
		{Name: "create", Usage: "create db", Action: doCreate},
		{Name: "drop", Usage: "drop db", Action: doDrop},
		{Name: "migrate", Usage: "migrate db", Action: doMigrate},
	}
	app.Action = action
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func action(c *cli.Context) {
	_ = cli.ShowAppHelp(c)
}

func doCreate(c *cli.Context) {
	conf := mysql2.GetDBConfig()
	utilDB := mysql.MakeDBUtil(conf)
	log.QyLogger.Info("do create", zap.String("db name", conf.DbName))
	utilDB.CreateDB()
}

func doDrop(c *cli.Context) {
	conf := mysql2.GetDBConfig()
	if strings.Contains(conf.DbName, "prod") {
		log.QyLogger.Warn("can't drop prod db", zap.String("db name", conf.DbName))
		return
	}
	utilDB := mysql.MakeDBUtil(conf)
	log.QyLogger.Info("do drop", zap.String("db name", conf.DbName))
	utilDB.DropDB()
}

func doMigrate(c *cli.Context) {
	log.QyLogger.Info("do migrate")
	mysqlDB := mysql.MakeDB(mysql2.GetDBConfig()).GetDB()
	mysql2.Migrate(mysqlDB)
}


