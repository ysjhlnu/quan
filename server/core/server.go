package core

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"quan/asset"
	"quan/global"
	"quan/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 quan
	当前版本:V0.0.1
    QQ群：620176501
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)

	// 定时任务
	c := cron.New()
	_, _ = c.AddFunc("0 9 * * * *", func() {
		fmt.Println("定时任务更新aws主机")
		asset.AssetHostAwsUpdate()
	})
	c.Start()

	global.GVA_LOG.Error(s.ListenAndServe().Error())

}
