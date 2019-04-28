package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
)

func TestProjectDB_GetProByToken(t *testing.T) {
	defer gDB.ClearAllData()

	pDB := NewProjectDB(gDB.GetDB())
	assert.NoError(t, pDB.Create(model.Project{
		ID: "p1",
		Name: "p1",
		Token: "ok",
		// 标志是否可以创建 notify
		Locked: false,

		UserID: "u1",

		CreatedAt: time.Now(),
	}))

	rp := pDB.GetProByToken("ok")
	assert.Equal(t, "p1", rp.ID)
	rp = pDB.GetProByToken("ok1")
	assert.Equal(t, "", rp.ID)
}
