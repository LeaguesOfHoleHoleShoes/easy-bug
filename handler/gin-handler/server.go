package gin_handler

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/g-env"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/common"
	"github.com/gin-gonic/gin"
)

func NewServer(h *common.Handler) *gin.Engine {
	if g_env.GetRunEnv() == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	// 如果有新的执行方法，又要兼容之前的，则加一个v2
	rgV1 := e.Group("v1")
	// 创建用户
	rgV1.POST("user", CreateUser(h))
	// 新增 pipeline （pipeline中包含了其对应的规则）
	rgV1.POST("login", Login(h))
	rgV1.POST("project", CreateProject(h))
	rgV1.POST("notify", CreateNotify(h))
	rgV1.GET("notifies", GetNotifies(h))
	rgV1.GET("notifies/latest", LatestNotifies(h))

	return e
}
