package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/c"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/alice52/archive/bili/util"
	"github.com/gookit/goutil/jsonutil"
	"github.com/wordpress-plus/kit-common/kg"
	"go.uber.org/zap"
)

type UserHistoryServiceIn struct{}

func (ce *UserHistoryServiceIn) SyncUserHistory() (err error) {
	items, err := api.LogonClient.UserHistory()
	if err != nil {
		return err
	}

	var vs []*model.ArchivedVideo
	for _, item := range items.Data.List {
		m := &model.ArchivedViewHistory{
			Bvid: item.History.Bvid,

			Title:     util.ToPrt(item.Title),
			Cover:     util.ToPrt(item.Cover),
			UpperMid:  item.AuthorMid,
			UpperName: item.AuthorName,
			FaceName:  item.AuthorFace,
			Duration:  item.Duration,
			ViewAt:    item.ViewAt,
			Videos:    item.Videos,
			Progress:  item.Progress,
			IsFinish:  item.IsFinish,
			Kid:       util.ToPrt(item.Kid),
			TagName:   util.ToPrt(item.TagName),
			Resp:      util.ToPrt(jsonutil.MustString(item)),
		}

		if err = dal.Q.ArchivedViewHistory.Save(m); err != nil {
			kg.L.Error("sync user history error", zap.Error(err))
		}
		if len(m.Bvid) > 0 {
			vs = append(vs, &model.ArchivedVideo{
				Bvid:         m.Bvid,
				ArchivedType: c.ArchivedTypeView})
		}
	}

	_ = UserVideoService.Merge(vs)
	return err
}
