package scheduler

import (
	"github.com/alice52/archive/bili/service"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/common/global"
	"github.com/micro-services-roadmap/kit-common/gormx/initialize"
	"github.com/micro-services-roadmap/kit-common/kg"
	"github.com/micro-services-roadmap/kit-common/viperx"
	"github.com/micro-services-roadmap/kit-common/zapx"
	"testing"
)

func init() {
	viperx.InitViper("../config-local.yaml")
	kg.L = zapx.Zap(global.CONFIG.Zap)

	// init db and do migration
	kg.DB = initialize.GormMysql(false)
	if kg.DB.Error != nil {
		panic(kg.DB.Error)
	} else {
		dal.SetDefault(kg.DB)
	}
}

func TestVideoFk(t *testing.T) {

	dal.SetDefault(kg.DB)
	find, err := dal.ArchivedCoin.Preload(dal.ArchivedCoin.ArchivedVideo).Find()
	if find != nil && err != nil {
		return
	}
}

func TestUserUpperTagService(t *testing.T) {
	err := service.UserUpperTagService.SyncUpperTags()
	if err != nil {
		panic(err)
	}
}

func TestSyncUppersService(t *testing.T) {
	err := service.UserUpperService.SyncUppers()
	if err != nil {
		panic(err)
	}
}

func TestSyncUserFavFolders(t *testing.T) {
	err := service.UserFavFolderService.SyncUserFavFolders()
	if err != nil {
		panic(err)
	}
}

func TestUserFavService(t *testing.T) {
	err := service.UserFavService.SyncUserFav()
	if err != nil {
		panic(err)
	}
}

func TestSyncUserLiked(t *testing.T) {

	for true {
		err := service.UserLikedService.SyncUserLiked()
		if err != nil {
			panic(err)
		}

	}
}

func TestSyncUserCoined(t *testing.T) {
	err := service.UserCoinedService.SyncUserCoined()
	if err != nil {
		panic(err)
	}
}

func TestSyncUserVideo(t *testing.T) {
	err := service.UserVideoService.SyncUserVideo()
	if err != nil {
		panic(err)
	}
}

func TestSyncUserHistory(t *testing.T) {
	err := service.UserHistoryService.SyncUserHistory()
	if err != nil {
		panic(err)
	}
}
