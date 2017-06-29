package example

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"ops-agent/core"
)

type ShortTask struct {
}

func (st *ShortTask) Run(ctx context.Context) error {
	fmt.Println("short task example")

	return nil
}

func (st *ShortTask) Name() string {
	return "short_task"
}

func (st *ShortTask) Kill() {
	return
}

func (st *ShortTask) Initialize(config *viper.Viper) error {
	return nil
}

func (st *ShortTask) Type() string {
	return core.TASK_SHORT_TERM
}
