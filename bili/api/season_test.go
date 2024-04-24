package api

import (
	"fmt"
	"github.com/alice52/archive/bili/util"
	"github.com/micro-services-roadmap/kit-common/viperx"
	"github.com/micro-services-roadmap/kit-common/zapx"
	"testing"
)

func init() {
	viperx.InitViper("../config-local.yaml")
	zapx.InitZap()
}

func TestSeasonSection(t *testing.T) {
	info, err := logonFunc().SeasonSection("", "729217") //
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
