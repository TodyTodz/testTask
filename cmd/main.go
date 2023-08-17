package main

import (
	"flag"

	"testTask/internal/app"
)

//Main function - for processing input data and run service.
//  In the future, there may be more of input data.
func main(){
	httpConfigPath := flag.String("http-config", "/etc/configs/http/settings.xml",
		"full path to config file")

	flag.Parse()

	app.Run(*httpConfigPath)

}
