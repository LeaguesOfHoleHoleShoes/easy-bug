package mysql

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetDBConfig(t *testing.T) {
	conf := GetDBConfig()
	assert.Equal(t, "easy_bug_dev", conf.DbName)
	assert.NoError(t, os.Setenv("db_env", "prod"))
	conf = GetDBConfig()
	assert.Equal(t, "easy_bug_prod", conf.DbName)
}
