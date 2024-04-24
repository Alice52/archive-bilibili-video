package api

import (
	"fmt"
	c "github.com/alice52/archive/bili/api/common"
	"github.com/alice52/archive/bili/util"
	"github.com/alice52/archive/common/global"
	"github.com/micro-services-roadmap/kit-common/kg"
	"github.com/micro-services-roadmap/kit-common/viperx"
	"github.com/micro-services-roadmap/kit-common/zapx"
	"testing"
)

func init() {
	viperx.InitViper("../config-local.yaml")
	kg.L = zapx.Zap(global.CONFIG.Zap)
}

func TestPlayUrl(t *testing.T) {
	info, err := logonFunc().PlayUrl("BV1CA411S7q4", 1031064040, c.Qn4k, c.FnvalHDR|c.Fnval4K)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
