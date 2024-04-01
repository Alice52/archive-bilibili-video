package api

import (
	"github.com/alice52/archive/common/global"
	"github.com/wordpress-plus/kit-common/kg"
	"github.com/wordpress-plus/kit-common/viperx"
	"github.com/wordpress-plus/kit-common/zapx"
)

func init() {
	viperx.InitViper("../config-local.yaml")
	kg.L = zapx.Zap(global.CONFIG.Zap)
}
