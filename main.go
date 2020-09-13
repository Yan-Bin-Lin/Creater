package main

import (
	"app/router"
	"app/setting"
	"log"
	"net/http"

	"fmt"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	runServe("main", router.HostSwitch{Engine: router.MainRouter()})
	runServe("account", router.HostSwitch{Engine: router.AccountRouter()})
	runServe("asset", router.HostSwitch{Engine: router.AssetRouter()})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func runServe(serve string, hs router.HostSwitch)  {
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.Servers[serve].Port),
		Handler:      hs,
		ReadTimeout:  setting.Servers[serve].ReadTimeout,
		WriteTimeout: setting.Servers[serve].WriteTimeout,
	}
	g.Go(func() error {
		return s.ListenAndServe()
	})
}