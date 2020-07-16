package main

import (
	"github.com/shadracnicholas/home-automation/libraries/go/bootstrap"
	"github.com/shadracnicholas/home-automation/service.user/handler"
)

//go:generate jrpc user.def

func main() {
	svc := bootstrap.Init(&bootstrap.Opts{
		ServiceName: "service.user",
		Database:    true,
	})

	r := handler.NewRouter(&handler.Controller{})
	svc.Run(r)
}
