package main

import (
	"github.com/shadracnicholas/home-automation/libraries/go/bootstrap"
	"github.com/shadracnicholas/home-automation/libraries/go/firehose"
	"github.com/shadracnicholas/home-automation/service.scene/consumer"
	"github.com/shadracnicholas/home-automation/service.scene/handler"
)

//go:generate jrpc scene.def

func main() {
	svc := bootstrap.Init(&bootstrap.Opts{
		ServiceName: "service.scene",
		Firehose:    true,
		Database:    true,
	})

	firehose.Subscribe(consumer.HandleSetSceneEvent)

	r := handler.NewRouter(&handler.Controller{})

	svc.Run(r)
}
