package main

import (
	"github.com/shadracnicholas/home-automation/libraries/go/bootstrap"
	"github.com/shadracnicholas/home-automation/libraries/go/router"
)

func main() {
	svc := bootstrap.Init(&bootstrap.Opts{
		ServiceName: "service.ping",
	})

	// The router has a default ping handler defined
	// in: libraries/go/router/middleware.go
	r := router.New()
	svc.Run(r)
}
