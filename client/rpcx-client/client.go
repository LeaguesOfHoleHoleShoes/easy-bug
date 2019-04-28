package rpcx_client

import (
	"context"
	"errors"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	common2 "github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/common"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"go.uber.org/zap"
	"runtime"
	"time"
)

var (
	ErrProTokenIsEmpty = errors.New("pro token is empty")
	ErrShouldLoginFirst = errors.New("should login first")
)

// pro token 从环境变量中读，每个环境是不同的
func NewEasyBugRpcxClient(serverUrl string, proToken string) (*EasyBugRpcxClient, error) {
	if proToken == "" {
		return nil, ErrProTokenIsEmpty
	}
	d := client.NewPeer2PeerDiscovery("tcp@" + serverUrl, "")

	opt := client.DefaultOption
	opt.Heartbeat = true
	opt.HeartbeatInterval = 8 * time.Second
	opt.CompressType = protocol.Gzip

	return &EasyBugRpcxClient{
		serverUrl: serverUrl,
		proToken: proToken,
		xClient: client.NewXClient("Handler", client.Failtry, client.RandomSelect, d, opt),
	}, nil
}

type EasyBugRpcxClient struct {
	serverUrl string
	proToken string
	xClient client.XClient

	username string
	password string
	userToken string
}

func (c *EasyBugRpcxClient) SetProToken(pt string) {
	c.proToken = pt
}

func (c *EasyBugRpcxClient) NotifyTest(title string, content string, extraData interface{}) {
	c.notify(title, content, extraData, model.NotifyTest)
}

func (c *EasyBugRpcxClient) NotifyEvent(title string, content string, extraData interface{}) {
	c.notify(title, content, extraData, model.NotifyEvent)
}

func (c *EasyBugRpcxClient) NotifyError(title string, content string, extraData interface{}) {
	c.notify(title, content, extraData, model.NotifyError)
}

func (c *EasyBugRpcxClient) notify(title string, content string, extraData interface{}, nt model.NotifyType) {
	var resp uint
	err := c.xClient.Call(context.Background(), "CreateNotify", &common2.CreateNotifyReq{
		ProToken: c.proToken,
		Title: title,
		Content: content,
		ExtraData: util.StringifyJson(extraData),
		NType: nt,
		System: runtime.GOOS,
	}, &resp)
	if err != nil {
		log.QyLogger.Warn("send notify failed", zap.Error(err))
	}
}

func (c *EasyBugRpcxClient) Login(username, password string) (model.User, error) {
	var resp common2.LoginResp
	if err := c.xClient.Call(context.Background(), "CreateNotify", &common2.LoginReq{
		Username: username,
		Password: password,
	}, &resp); err != nil {
		return model.User{}, err
	}
	c.username = username
	c.password = password
	c.userToken = resp.User.Token
	return resp.User, nil
}

func (c *EasyBugRpcxClient) CreateProject(name string) error {
	if c.userToken == "" {
		return ErrShouldLoginFirst
	}
	var resp common2.CreateProjectResp
	return c.xClient.Call(context.Background(), "CreateProject", &common2.CreateProjectReq{
		UserToken: c.userToken,
		Name: name,
	}, &resp)
}

func (c *EasyBugRpcxClient) CreateUser(username, password, name, phone string) error {
	var resp uint
	return c.xClient.Call(context.Background(), "CreateUser", &common2.CreateUserReq{
		Username: username,
		Password: password,
		Name: name,
		Phone: phone,
	}, &resp)
}

func (c *EasyBugRpcxClient) GetProjects(page, perPage int) ([]model.Project, error) {
	if c.userToken == "" {
		return nil, ErrShouldLoginFirst
	}
	var resp common2.GetProjectsResp
	err := c.xClient.Call(context.Background(), "GetProjects", &common2.GetProjectsReq{
		UserToken: c.userToken,
		Page: page,
		PerPage: perPage,
	}, &resp)
	return resp.Projects, err
}

func (c *EasyBugRpcxClient) LatestNotifies(count int) ([]model.Notify, error) {
	var resp common2.LatestNotifiesResp
	err := c.xClient.Call(context.Background(), "LatestNotifies", &common2.LatestNotifiesReq{
		ProToken: c.proToken,
		Count: count,
	}, &resp)
	return resp.Notifies, err
}