package main

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/scheduler"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/common/global"
	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/kit-common"
	"github.com/micro-services-roadmap/kit-common/api/middleware"
	"github.com/micro-services-roadmap/kit-common/kg"
)

func init() {
	// init viper
	kit.InitWithConf(&global.CONFIG, "config-local.yaml")
	dal.SetDefault(kg.DB)
	go func() {
		api.LogonClient, _ = api.GetLogonClient()
	}()
}

func main() {
	scheduler.Scheduler()
	r := gin.Default()
	r.Use(middleware.GinRecovery(true))
	if err := r.Run(":8888"); err != nil {
		kg.L.Error(err.Error())
	}
}
