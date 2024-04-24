package core

import (
	"github.com/alice52/archive/common/util"
	"github.com/micro-services-roadmap/kit-common/kg"
	"testing"
)

func init() {
	kg.V = viperx.Viper(&kg.C, "config.yaml")
}

func TestEmail(_ *testing.T) {

	err := util.EmailTest("Test Email", "Test Body")
	if err != nil {
		return
	}
}
