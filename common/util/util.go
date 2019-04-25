package util

import (
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
