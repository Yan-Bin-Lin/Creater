package main

import (
	"app/router"
	"app/setting"

	"fmt"
	"log"
	"net/http"
)

var hs router.HostSwitch

func main() {
	//	accountRouter := router.AcountRouter()
	//	assetRouter := router.AssetRouter()

	// Make a new HostSwitch and insert the router (our http handler)
	hs = make(router.HostSwitch)
	hs[fmt.Sprintf("%s:%d", setting.Servers["main"].Host, setting.Servers["main"].Port)] = router.MainRouter()
	hs[fmt.Sprintf("%s:%d", setting.Servers["account"].Host, setting.Servers["account"].Port)] = router.AccountRouter()
	hs[fmt.Sprintf("%s:%d", setting.Servers["asset"].Host, setting.Servers["asset"].Port)] = router.AssetRouter()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.Servers["main"].Port),
		Handler:      hs,
		ReadTimeout:  setting.Servers["main"].ReadTimeout,
		WriteTimeout: setting.Servers["main"].WriteTimeout,
	}

	log.Fatal(s.ListenAndServe())
}
