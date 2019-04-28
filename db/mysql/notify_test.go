package mysql

import (
	"fmt"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/orm/gorm/mysql"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var gDB mysql.DB

func init() {
	if err := os.Setenv("db_env", "test"); err != nil {
		panic(err)
	}
	conf := GetDBConfig()

	mDB := mysql.MakeDBUtil(conf)
	mDB.DropDB()
	mDB.CreateDB()
	gDB = mysql.MakeDB(conf)
	Migrate(gDB.GetDB())
}

func TestNotifyDB_Create(t *testing.T) {
	defer gDB.ClearAllData()

	nDB := NewNotifyDB(gDB.GetDB())
	n1 := model.Notify{
		Title: "n1",
		Content: "hello",
		ExtraData: "ex",
		System: "linux",
		StackInfo: "s",
		NType: model.NotifyTest,

		ID: "123",
		FromUrl: "f u",
		FromIP: "ip",
		Region: "sz",

		ProjectID: "p1",
	}
	assert.NoError(t, nDB.Create(n1))
	assert.Error(t, nDB.Create(n1))

	n2 := n1
	n2.ID = "n2"
	n2.ProjectID = "p2"
	assert.NoError(t, nDB.Create(n2))

	for i := 3; i < 6; i++ {
		nx := n1
		nx.ID = fmt.Sprintf("n%v", i)
		assert.NoError(t, nDB.Create(nx))
	}

	p2R := nDB.LatestNotifies("p2", 3)
	assert.Len(t, p2R, 1)
	assert.Equal(t, "n2", p2R[0].ID)

	p1R, _, _ := nDB.Notifies("p1", 1, 2)
	assert.Len(t, p1R, 2)
}
