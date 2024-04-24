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

type UserCoinedServiceIn struct{}

func (ce *UserCoinedServiceIn) SyncUserCoined() (err error) {
	items, err := api.LogonClient.UserCoined()
	if err != nil {
		return err
	}

	var vs []*model.ArchivedVideo
	for _, item := range items.Data {
		m := &model.ArchivedCoin{
			Bvid:       item.Bvid,
			Aid:        item.Aid,
			Cid:        item.Cid,
			Coins:      item.Coins,
			Cover:      &item.Pic,
			Duration:   item.Duration,
			CoinedTime: item.Ctime,
			Intro:      &item.Desc,
			Title:      &item.Title,
			Type:       item.Tid,
		}

		up := jsonutil.MustString(item.Owner)
		m.Owner = &up
		cnt := jsonutil.MustString(item.Stat)
		m.CntInfo = &cnt
		resp := jsonutil.MustString(item)
		m.Resp = &resp

		if err = dal.Q.ArchivedCoin.Save(m); err != nil {
			kg.L.Error("sync user coined error", zap.Error(err))
		}

		if len(item.Bvid) > 0 {
			vs = append(vs, &model.ArchivedVideo{
				Bvid:         item.Bvid,
				ArchivedType: c.ArchivedTypeCoin})
		}
	}
	_ = UserVideoService.Merge(vs)
	return err
}
