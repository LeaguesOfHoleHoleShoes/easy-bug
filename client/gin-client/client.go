package gin_client

import (
	"bytes"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/common/util"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/common"
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/model"
	"github.com/dipperin/go-ms-toolkit/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"runtime"
)

func NewEasyBugGinClient(serverUrl string, proToken string) *EasyBugGinClient {
	return &EasyBugGinClient{serverUrl: serverUrl, proToken: proToken}
}

type EasyBugGinClient struct {
	serverUrl string
	proToken string
}

func (c *EasyBugGinClient) SetProToken(pt string) {
	c.proToken = pt
}

func (c *EasyBugGinClient) NotifyTest(title string, content string, extraData interface{}) {
	c.notify(title, content, extraData, model.NotifyTest)
}

func (c *EasyBugGinClient) NotifyEvent(title string, content string, extraData interface{}) {
	c.notify(title, content, extraData, model.NotifyEvent)
}

func (c *EasyBugGinClient) NotifyError(title string, content string, extraData interface{}) {
	c.notify(title, content, extraData, model.NotifyError)
}

func (c *EasyBugGinClient) notify(title string, content string, extraData interface{}, nt model.NotifyType) {
	resp, err := http.Post(c.serverUrl + "/v1/notify", "application/json", bytes.NewReader(util.StringifyJsonToBytes(common.CreateNotifyReq{
		ProToken: c.proToken,
		Title: title,
		Content: content,
		ExtraData: util.StringifyJson(extraData),
		NType: nt,
		System: runtime.GOOS,
	})))
	if err != nil {
		log.QyLogger.Warn("send notify failed", zap.Error(err))
		return
	}
	if resp.StatusCode == 200 {
		return
	}
	hrb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.QyLogger.Warn("read resp body failed", zap.Error(err))
		return
	}

	var respData common.BaseResp
	if err = util.ParseJsonFromBytes(hrb, &respData); err != nil {
		log.QyLogger.Warn("parse notify result failed", zap.Error(err))
	}

	if !respData.Success {
		log.QyLogger.Error("send notify failed", zap.String("info", respData.Info))
	}
}