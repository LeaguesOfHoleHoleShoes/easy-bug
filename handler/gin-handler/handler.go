package gin_handler

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/common"
	"github.com/gin-gonic/gin"
)

func CreateUser(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.CreateUserReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		var resp uint
		if err := h.CreateUser(context, &req, &resp); err != nil {
			errResp(context, err.Error())
			return
		}
		successResp(context)
	}
}

func Login(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.LoginReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		var resp common.LoginResp
		_ = h.Login(context, &req, &resp)
		doResp(context, resp.BaseResp, resp)
	}
}

func CreateProject(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.CreateProjectReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		var resp common.CreateProjectResp
		_ = h.CreateProject(context, &req, &resp)
		doResp(context, resp.BaseResp, resp)
	}
}

func CreateNotify(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.CreateNotifyReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		// todo FromUrl FromIP Region
		var resp uint
		if err := h.CreateNotify(context, &req, &resp); err != nil {
			errResp(context, err.Error())
			return
		}
		successResp(context)
	}
}

func LatestNotifies(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.LatestNotifiesReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		var resp common.LatestNotifiesResp
		_ = h.LatestNotifies(context, &req, &resp)
		doResp(context, resp.BaseResp, resp)
	}
}

func GetNotifies(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.GetNotifiesReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		var resp common.GetNotifiesResp
		_ = h.GetNotifies(context, &req, &resp)
		doResp(context, resp.BaseResp, resp)
	}
}

func GetProjects(h *common.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req common.GetProjectsReq
		if err := context.BindJSON(&req); err != nil {
			errResp(context, err.Error())
			return
		}
		var resp common.GetProjectsResp
		_ = h.GetProjects(context, &req, &resp)
		doResp(context, resp.BaseResp, resp)
	}
}

func errResp(context *gin.Context, err string) {
	context.JSON(400, common.BaseResp{
		Success: false,
		Info: err,
	})
}

func successResp(context *gin.Context) {
	context.JSON(200, common.BaseResp{
		Success: true,
	})
}

func doResp(context *gin.Context, baseR common.BaseResp, data interface{}) {
	if baseR.Success {
		context.JSON(200, data)
	} else {
		context.JSON(400, data)
	}
}