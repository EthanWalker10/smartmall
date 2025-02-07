package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/EthanWalker10/smartmall/api/user-web/utils/register/consul"
	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/EthanWalker10/smartmall/api/user-web/global"
	"github.com/EthanWalker10/smartmall/api/user-web/initialize"
	"github.com/EthanWalker10/smartmall/api/user-web/utils"
	myvalidator "github.com/EthanWalker10/smartmall/api/user-web/validator"
)

func main() {
	initialize.InitLogger()

	//2. 初始化配置文件
	initialize.InitConfig()

	Router := initialize.Routers()

	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//5. 初始化srv的连接
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境启动获取端口号
	debug := viper.GetBool("MXSHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	// register gin’s validator of mobile
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// register custom validator for 'mobile' tag
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		// register translation for 'mobile' tag
		// cause the default translator doesn’t exert an effext in the validator we register above
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//服务注册
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	/*
		1. S() gets a global sugar
		2. S() and L() are very useful, providing a global safe access logger
	*/
	zap.S().Debugf("Server is starting, port： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("Failed to start the server:", err.Error())
	}
	zap.S().Infof("Server is running on: http://%s:%d", global.ServerConfig.Host, global.ServerConfig.Port)


	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//if err = register_client.DeRegister(serviceId); err != nil {
	//	zap.S().Info("注销失败:", err.Error())
	//}else{
	//	zap.S().Info("注销成功:")
	//}
}
