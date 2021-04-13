package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/librespeed/speedtest/config"
	"github.com/librespeed/speedtest/web"
)

var (
	optConfig = flag.String("c", "", "config file to be used, defaults to settings.toml in the same directory")
)

func main() {
	flag.Parse()
	conf := config.Load(*optConfig)
	web.SetServerLocation(&conf)
	log.Fatal(web.ListenAndServe(&conf))
}
