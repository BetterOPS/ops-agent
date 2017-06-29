package core

import (
	"context"
	"github.com/spf13/viper"
)

type Task interface {
	Run(context.Context) error
	Kill()
	Name() string
	Type() string
	Initialize(config *viper.Viper) error
}
