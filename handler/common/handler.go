package common

import (
	"context"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
)

func NewHandler(env Env) *Handler {
	return &Handler{Env: env}
}

type Handler struct {
	Env
}

func (h *Handler) CreateUser(c context.Context, req *CreateUserReq, resp *uint) error {
	return h.uManager.Create(model.User{
		Username: req.Username,
		Password: req.Password,
		Name: req.Name,
		Phone: req.Phone,
	})
}

func (h *Handler) Login(c context.Context, req *LoginReq, resp *LoginResp) error {
	user, err := h.uManager.Login(req.Username, req.Password)
	if err != nil {
		errResp(&resp.BaseResp, err.Error())
		return err
	}
	resp.User = user
	resp.Success = true

	return nil
}

func (h *Handler) CreateProject(c context.Context, req *CreateProjectReq, resp *CreateProjectResp) error {
	user, err := h.uManager.ValidUserToken(req.UserToken)
	if err != nil {
		errResp(&resp.BaseResp, err.Error())
		return err
	}

	resultP, err := h.pManager.Create(model.Project{
		Name: req.Name,
		UserID: user.ID,
	})
	if err != nil {
		errResp(&resp.BaseResp, err.Error())
		return err
	}

	resp.Project = resultP
	resp.Success = true

	return nil
}

/*

外边赋值
FromUrl string `json:"from_url"`
FromIP string `json:"from_ip"`
Region string `json:"region"`

*/
func (h *Handler) CreateNotify(c context.Context, req *CreateNotifyReq, resp *uint) error {
	return h.nManager.Create(req.ProToken, req.GetNotify())
}

func (h *Handler) LatestNotifies(c context.Context, req *LatestNotifiesReq, resp *LatestNotifiesResp) error {
	ns, err := h.nManager.LatestNotifies(req.ProToken, req.Count)
	if err != nil {
		errResp(&resp.BaseResp, err.Error())
		return err
	}
	resp.Notifies = ns
	resp.Success = true
	return nil
}

func (h *Handler) GetNotifies(c context.Context, req *GetNotifiesReq, resp *GetNotifiesResp) (err error) {
	resp.Notifies, resp.TotalPages, resp.TotalCount, err = h.nManager.Notifies(req.ProToken, req.Page, req.PerPage)
	if err != nil {
		errResp(&resp.BaseResp, err.Error())
		return
	}
	resp.Success = true
	return
}

func errResp(br *BaseResp, err string) {
	br.Success = false
	br.Info = err
}