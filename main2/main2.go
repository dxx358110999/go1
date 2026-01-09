package main2

import (
	"dxxproject/config_prepare/app_config"
	"dxxproject/ioc"
	"dxxproject/pkg/nacos_ok"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Main2() (err error) {
	injector, err := ioc.Inject()
	if err != nil {
		return err
	}
	nacosInstance := do.MustInvoke[*nacos_ok.NacosInstance](injector)
	appConfig := do.MustInvoke[*app_config.AppConfig](injector)
	engine := do.MustInvoke[*gin.Engine](injector)

	redClient := do.MustInvoke[*redis.Client](injector) //退出时关闭连接
	defer redClient.Close()

	if err != nil {
		return err
	}

	go func() {
		err = engine.Run(fmt.Sprintf(":%d", appConfig.Port))
	}()
	if err != nil {
		fmt.Printf("gin启动失败:%v\n", err)
		os.Exit(1)
	}

	time.Sleep(3 * time.Second)
	err = nacosInstance.Register()
	if err != nil {
		fmt.Printf("注册:%v\n", err)
		return
	}

	// 等待中断信号以优雅地退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 取消服务注册
	err = nacosInstance.Deregister()
	if err != nil {
		log.Fatalln("反注册失败:", err)
	}

	//等待通知结束
	if appConfig.Mode == "release" {
		time.Sleep(15 * time.Second)
	} else {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("主线程退出")
	return
}
