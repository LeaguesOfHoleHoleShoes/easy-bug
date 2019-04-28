package common

import "github.com/LeaguesOfHoleHoleShoes/easy-bug/model"

type BaseResp struct {
	Success bool `json:"success"`
	Info string `json:"info"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`

	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	BaseResp
	User model.User `json:"user"`
}

type CreateProjectReq struct {
	Name string `json:"name"`
	UserToken string `json:"user_token"`
}

type CreateProjectResp struct {
	BaseResp
	Project model.Project `json:"project"`
}

type CreateNotifyReq struct {
	ProToken string `json:"pro_token"`

	Title string `json:"title"`
	Content string `json:"content"`
	ExtraData string `json:"extra_data"`
	System string `json:"system"`
	StackInfo string `json:"stack_info"`
	// test event error
	NType model.NotifyType `json:"n_type"`

	FromUrl string `json:"from_url"`
	FromIP string `json:"from_ip"`
	Region string `json:"region"`
}

func (r CreateNotifyReq) GetNotify() model.Notify {
	return model.Notify{
		Title: r.Title,
		Content: r.Content,
		ExtraData: r.ExtraData,
		System: r.System,
		StackInfo: r.StackInfo,
		// test event error
		NType: r.NType,

		FromUrl: r.FromUrl,
		FromIP: r.FromIP,
		Region: r.Region,
	}
}

type LatestNotifiesReq struct {
	ProToken string `json:"pro_token"`
	Count int `json:"count"`
}

type LatestNotifiesResp struct {
	BaseResp
	Notifies []model.Notify `json:"notifies"`
}

type GetNotifiesReq struct {
	ProToken string `json:"pro_token"`
	Page int `json:"page"`
	PerPage int `json:"per_page"`
}

type GetNotifiesResp struct {
	BaseResp
	Notifies []model.Notify `json:"notifies"`
	TotalPages int `json:"total_pages"`
	TotalCount int `json:"total_count"`
}