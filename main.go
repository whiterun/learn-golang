package main

import (
	"flag"
	"bagogo/modules/api"
	"bagogo/modules/util/config"
)

func main() {
	// Load environment configuration from file
	cfgPath := flag.String("p", "./config.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
