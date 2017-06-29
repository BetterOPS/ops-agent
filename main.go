package main

import (
	"context"
	"github.com/spf13/viper"
	"ops-agent/core"
	"ops-agent/example"
	"ops-agent/lib/log"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
)

func init() {
	viper.SetConfigName("agent")
	viper.AddConfigPath("/etc/ops-agent")
	viper.AddConfigPath("$HOME/.ops-agent")
	viper.AddConfigPath("./etc")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.InitLogger(viper.Sub("log").GetString("level"))
	go func() {
		sup := core.NewSupervisor(ctx)
		defer sup.Stop()
		// add task
		sup.Run(&example.ShortTask{})
		sup.Start()
	}()
	InitSignal()
}
