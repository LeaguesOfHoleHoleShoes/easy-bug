package util

import (
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"os"
	"os/user"
	"path/filepath"
)

// 解析json字符串
func ParseJson(data string, result interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal([]byte(data), result)
}

// json转字符串
func StringifyJson(data interface{}) string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _  := json.Marshal(&data)
	return string(b)
}

// 解析json bytes
func ParseJsonFromBytes(data []byte, result interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, result)
}

// json bytes转字符串
func StringifyJsonToBytes(data interface{}) []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.Marshal(&data)
	return b
}

func StringifyJsonToBytesWithErr(data interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(&data)
	return b, err
}

func GopathDir() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return filepath.Join(HomeDir(), "go")
	}
	return gopath
}

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

// 获取offset
func GetOffset(page int, perPage int) int {
	if page < 1 {
		page = 1
	}
	return (page - 1) * perPage
}

// 获取分页页数总数及数据列表。m是查询的表的model，result是列表结果传指针进来！
func GetDataByPageAndPerPage(db *gorm.DB, page int, perPage int, m interface{}, result interface{}) (totalPages int, totalCount int) {
	offset := GetOffset(page, perPage)
	if err := db.Offset(offset).Limit(perPage).Find(result).Error; err != nil {
		return
	}

	db = db.Offset(-1).Limit(-1)
	if err := db.Model(m).Count(&totalCount).Error; err != nil {
		return
	}
	if perPage <= 0 {
		return
	}
	totalPages = (totalCount + perPage -1)/perPage
	return
}
