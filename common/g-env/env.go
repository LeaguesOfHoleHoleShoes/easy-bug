package g_env

import (
	"github.com/dipperin/go-ms-toolkit/log"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

// 处理 docker 中 log 到 /var/log/qy 中
func InitLogger() {
	switch GetDockerEnv() {
	case "", "0":
	default:
		dockerLogDir := filepath.Join("/var/log/hh", "easy-bug")
		if err := os.MkdirAll(dockerLogDir, 0755); err != nil {
			panic(err)
		}
		podName := os.Getenv("HOSTNAME")
		if podName == "" {
			panic("can't get HOSTNAME from env")
		}
		log.InitLogger(zapcore.DebugLevel, dockerLogDir, podName + ".log", false)
	}
}

// 0 非docker环境 1 docker中非生产环境 2 docker中生产环境
func GetDockerEnv() string {
	return os.Getenv("docker_env")
}

// db名称配置 dev test preprod prod
func GetDBEnv() string {
	return os.Getenv("db_env")
}

// 程序执行环境配置 dev test preprod prod
func GetRunEnv() string {
	return os.Getenv("run_env")
}
