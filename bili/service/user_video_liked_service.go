package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/c"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/gookit/goutil/jsonutil"
	"github.com/micro-services-roadmap/kit-common/kg"
	"go.uber.org/zap"
)

type UserLikedServiceIn struct{}

func (ce *UserLikedServiceIn) SyncUserLiked() (err error) {
	items, err := api.LogonClient.UserLiked()
	if err != nil {
		return err
	}

	var vs []*model.ArchivedVideo
	for _, item := range items.Data.List {
		m := &model.ArchivedLike{
			Bvid:     item.Bvid,
			Aid:      item.Aid,
			Cid:      item.Cid,
			Cover:    &item.Pic,
			Duration: item.Duration,
			LikeTime: item.Ctime,
			SeasonID: item.SeasonID,
			Intro:    &item.Desc,
			Title:    &item.Title,
			Type:     item.Tid,
		}

		up := jsonutil.MustString(item.Owner)
		m.Owner = &up
		cnt := jsonutil.MustString(item.Stat)
		m.CntInfo = &cnt
		resp := jsonutil.MustString(item)
		m.Resp = &resp

		if err = dal.Q.ArchivedLike.Save(m); err != nil {
			kg.L.Error("sync user liked error", zap.Error(err))
		}
		if len(item.Bvid) > 0 {
			vs = append(vs, &model.ArchivedVideo{
				Bvid:         item.Bvid,
				ArchivedType: c.ArchivedTypeLike})
		}
	}

	_ = UserVideoService.Merge(vs)
	return err
}
