package scheduler

import (
	"fmt"
	"github.com/alice52/archive/bili/service"
	"github.com/alice52/archive/common/global"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log"
)

func Scheduler() {
	// 创建一个新的Cron调度器, 使用 Zap 做日志记录
	stdOutLogger := zap.NewStdLog(global.LOG)
	logger := cron.VerbosePrintfLogger(log.New(stdOutLogger.Writer(), "cron: ", log.Llongfile))
	c := cron.New(cron.WithSeconds(), cron.WithLogger(logger))

	// 添加一个函数到调度器中，使其在指定的时间执行
	c.AddFunc("0 * * * * *", func() {
		fmt.Println("执行定时任务")
		global.LOG.Info("执行定时任务")
	})

	c.AddFunc("0 0 0/4 * * ?", func() {
		if err := service.UserUpperTagService.SyncUpperTags(); err != nil {
			global.LOG.Error("同步UP主标签失败", zap.Error(err))
		}
		global.LOG.Info("同步UP主标签成功")
	})

	c.AddFunc("0 1 0/4 * * ?", func() {
		if err := service.UserUpperService.SyncUppers(); err != nil {
			global.LOG.Error("同步UP主失败", zap.Error(err))
		}
		global.LOG.Info("同步UP主成功")
	})

	c.AddFunc("0 2 0/4 * * ?", func() {
		if err := service.UserFavFolderService.SyncUserFavFolders(); err != nil {
			global.LOG.Error("同步获取指定用户创建的所有收藏夹信息失败", zap.Error(err))
		}
		global.LOG.Info("同步获取指定用户创建的所有收藏夹信息成功")
	})

	c.AddFunc("0 3 0/4 * * ?", func() {
		if err := service.UserFavService.SyncUserFav(); err != nil {
			global.LOG.Error("同步获取收藏夹内容明细列表失败", zap.Error(err))
		}
		global.LOG.Info("同步获取收藏夹内容明细列表成功")
	})

	c.AddFunc("0 1 0/1 * * ?", func() {
		if err := service.UserLikedService.SyncUserLiked(); err != nil {
			global.LOG.Error("同步查询用户最近点赞视频失败", zap.Error(err))
		}
		global.LOG.Info("同步查询用户最近点赞视频成功")
	})
	c.AddFunc("0 0 0/1 * * ?", func() {
		if err := service.UserCoinedService.SyncUserCoined(); err != nil {
			global.LOG.Error("同步查询用户最近投币视频失败", zap.Error(err))
		}
		global.LOG.Info("同步查询用户最近投币视频成功")
	})

	// 启动调度器
	c.Start()
}
