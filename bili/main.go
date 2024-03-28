package main

import (
	"github.com/alice52/archive/bili/scheduler"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/common/global"
	initialize "github.com/alice52/archive/common/init"
	"github.com/gin-gonic/gin"
	"github.com/wordpress-plus/kit-logger/viperx"
	"github.com/wordpress-plus/kit-logger/zapx"
	"go.uber.org/zap"
)

func main() {
	// init viper
	global.VIPER = viperx.Viper(&global.CONFIG, "config-local.yaml") // 初始化Viper
	global.LOG = zapx.Zap(global.CONFIG.Zap)
	zap.ReplaceGlobals(global.LOG)

	// init db and do migration
	global.DB = initialize.GormMysql(true)
	if global.DB.Error != nil {
		panic(global.DB.Error)
	} else {
		dal.SetDefault(global.DB)
	}

	scheduler.Scheduler()
	r := gin.Default()
	if err := r.Run(":8888"); err != nil {
		global.LOG.Error(err.Error())
	}
}
