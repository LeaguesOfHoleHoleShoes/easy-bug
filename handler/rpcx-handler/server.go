package rpcx_handler

import (
	"github.com/LeaguesOfHoleHoleShoes/easy-bug/handler/common"
	"github.com/smallnest/rpcx/server"
)

func NewServer(h *common.Handler) *server.Server {
	s := server.NewServer()
	err := s.RegisterName("Handler", h, "")
	if err != nil {
		panic(err)
	}
	return s
}
