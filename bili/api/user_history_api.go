package api

import (
	"encoding/json"
	"fmt"
	m "github.com/alice52/archive/bili/api/model"
)

var (
	historyUrl = "https://api.bilibili.com/x/web-interface/history/cursor?type=archive&ps=20"
)

// region response model

type Histories struct {
	Cursor Cursor `json:"cursor"`
	List   []List `json:"list"`
	Tab    []Tab  `json:"tab"`
}

type Cursor struct {
	Business string `json:"business"`
	Max      int64  `json:"max"`
	PS       int64  `json:"ps"`
	ViewAt   int64  `json:"view_at"`
}

type List struct {
	AuthorFace string      `json:"author_face"`
	AuthorMid  int64       `json:"author_mid"`
	AuthorName string      `json:"author_name"`
	Badge      string      `json:"badge"`
	Cover      string      `json:"cover"`
	Covers     interface{} `json:"covers"`
	Current    string      `json:"current"`
	Duration   int64       `json:"duration"`
	History    History     `json:"history"`
	IsFav      int64       `json:"is_fav"`
	IsFinish   int64       `json:"is_finish"`
	Kid        int64       `json:"kid"`
	LiveStatus int64       `json:"live_status"`
	LongTitle  string      `json:"long_title"`
	NewDesc    string      `json:"new_desc"`
	Progress   int64       `json:"progress"`
	ShowTitle  string      `json:"show_title"`
	TagName    string      `json:"tag_name"`
	Title      string      `json:"title"`
	Total      int64       `json:"total"`
	URI        string      `json:"uri"`
	Videos     int64       `json:"videos"`
	ViewAt     int64       `json:"view_at"`
}

type History struct {
	Business string `json:"business"`
	Bvid     string `json:"bvid"`
	Cid      int64  `json:"cid"`
	Dt       int64  `json:"dt"`
	Epid     int64  `json:"epid"`
	OID      int64  `json:"oid"`
	Page     int64  `json:"page"`
	Part     string `json:"part"`
}

type Tab struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//endregion

// UserHistory 获取当前登录用户当天的历史记录
// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/history&toview/history.md#%E8%8E%B7%E5%8F%96%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95%E5%88%97%E8%A1%A8_web%E7%AB%AF
func (client *BClient) UserHistory() (*m.BResp[Histories], error) {

	info := &m.BResp[Histories]{}
	if ss, err := client.Get(fmt.Sprintf(historyUrl)); err != nil {
		return nil, err
	} else {
		return info, json.Unmarshal(ss, &info)
	}
}
