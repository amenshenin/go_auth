package main

import (
	"flag"
	"fmt"

	config "github.com/amenshenin/go_auth"
)

func main() {
	//Init config
	configPath := flag.String("config-path", "/run/secrets/config", "Please set in the config-path flag the path to config file")
	flag.Parse()
	config := config.MustLoad(*configPath)

	//Init logs

	//Init DB

	//Init server

	fmt.Println("go-go-go", config) //https://www.youtube.com/watch?v=rCJvW2xgnk0
}
