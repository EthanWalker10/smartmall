package global

import (
	"github.com/EthanWalker10/smartmall/api/user_web/config"
	"github.com/EthanWalker10/smartmall/api/user_web/proto"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	UserSrvClient proto.UserClient
)
