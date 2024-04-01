package api

import (
	"fmt"
	"github.com/alice52/archive/bili/util"
	"github.com/alice52/archive/common/global"
	"github.com/wordpress-plus/kit-common/kg"
	"github.com/wordpress-plus/kit-common/viperx"
	"github.com/wordpress-plus/kit-common/zapx"
	"testing"
)

func init() {
	viperx.InitViper("../config-local.yaml")
	kg.L = zapx.Zap(global.CONFIG.Zap)
}

func TestUserLiked(t *testing.T) {
	info, err := logonFunc().UserLiked()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
