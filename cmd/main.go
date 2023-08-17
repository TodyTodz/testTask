package main

import (
	"flag"

	"testTask/internal/app"
)

func main(){
	httpConfigPath := flag.String("http-config", "/etc/multicheck/configs/http/config.xml",
		"full path to config file")

	flag.Parse()

	app.Run(*httpConfigPath)

}
