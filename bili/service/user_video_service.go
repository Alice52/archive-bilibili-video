package service

import (
	"github.com/alice52/archive/bili/api"
	"github.com/alice52/archive/bili/c"
	"github.com/alice52/archive/bili/source/gen/dal"
	"github.com/alice52/archive/bili/source/gen/model"
	"github.com/gookit/goutil/jsonutil"
	"github.com/micro-services-roadmap/kit-common/kg"
	"github.com/micro-services-roadmap/kit-common/util"
	"go.uber.org/zap"
	"time"
)

type UserVideoServiceIn struct{}

func (ce *UserVideoServiceIn) Merge(vs []*model.ArchivedVideo) (err error) {

	if len(vs) == 0 {
		return
	}

	for _, v := range vs {
		v.CreateTime = util.ToPrt(time.Now())
		if err = dal.Q.ArchivedVideo.Create(v); err != nil {
			kg.L.Error("create video error", zap.Error(err))
		}
	}

	return
}

func (ce *UserVideoServiceIn) SyncUserVideo() (err error) {
	videos, err := dal.Q.ArchivedVideo.
		Select(dal.Q.ArchivedVideo.Bvid, dal.Q.ArchivedVideo.ArchivedType).
		Where(dal.Q.ArchivedVideo.SyncStatus.Eq(c.SyncStatusTodo)).Find()
	if err != nil {
		return err
	}

	for _, video := range videos {
		if err = DoSyncVideo(video.Bvid, video.ArchivedType); err != nil {
			kg.L.Error("sync video"+video.Bvid+" error", zap.Error(err))
		}
	}

	return err
}

func DoSyncVideo(bvid string, atype int64) (err error) {
	infos, err := api.LogonClient.VideoInfo(bvid)
	if err != nil {
		return err
	}

	r := infos.Data
	m := &model.ArchivedVideo{
		ArchivedType: atype,
		Bvid:         bvid,
		SyncStatus:   c.SyncStatusDoing,
		//SyncTime:     util.ToPrt(time.Now()),
		// SeasonID: resp.sea,
		Aid:            int64(r.Aid),
		Cid:            int64(r.Cid),
		Tid:            int64(r.Tid),
		Cover:          util.ToPrt(r.Pic),
		Ctime:          int64(r.Ctime),
		Pubdate:        int64(r.Pubdate),
		Duration:       int64(r.Duration),
		Title:          util.ToPrt(r.Title),
		Intro:          util.ToPrt(r.Desc),
		UpperMid:       int64(r.Owner.Mid),
		UpperName:      r.Owner.Name,
		FaceName:       r.Owner.Face,
		StatView:       int64(r.Stat.View),
		StatDanmaku:    int64(r.Stat.Danmaku),
		StatReply:      int64(r.Stat.Reply),
		StatFavorite:   int64(r.Stat.Favorite),
		StatCoin:       int64(r.Stat.Coin),
		StatShare:      int64(r.Stat.Share),
		StatLike:       int64(r.Stat.Like),
		StatDislike:    int64(r.Stat.Dislike),
		StatNowRank:    int64(r.Stat.Dislike),
		StatHisRank:    int64(r.Stat.Dislike),
		StatEvaluation: util.ToPrt(r.Stat.Evaluation),
		HonorReply:     util.ToPrt(jsonutil.MustString(r.HonorReply)),
		Resp:           util.ToPrt(jsonutil.MustString(infos)),
	}

	if err = dal.Q.ArchivedVideo.Save(m); err != nil {
		kg.L.Error("sync video error", zap.Error(err))
	}

	return err
}
