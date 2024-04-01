package main

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/scheduler"
	"github.com/gin-gonic/gin"
	"github.com/wordpress-plus/kit-common"
	"github.com/wordpress-plus/kit-common/kg"
)

func init() {
	// init viper
	kit.Init("config-local.yaml")
	go func() {
		api.LogonClient, _ = api.GetLogonClient()
	}()
}

func main() {
	scheduler.Scheduler()
	r := gin.Default()
	if err := r.Run(":8888"); err != nil {
		kg.L.Error(err.Error())
	}
}
