package main

import (
	"flag"

	"testTask/internal/app"
)

func main(){
	httpConfigPath := flag.String("http-config", "/etc/configs/http/settings.xml",
		"full path to config file")

	flag.Parse()

	app.Run(*httpConfigPath)

}
