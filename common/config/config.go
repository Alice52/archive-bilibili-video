package config

import (
	"github.com/micro-services-roadmap/kit-common/viperx/vconfig"
)

type Server struct {
	vconfig.Server `yaml:",inline" mapstructure:",squash"`
	Email          Email `mapstructure:"email" json:"email" yaml:"email"`
}
