package main

import (
	"flag"
	"fmt"
)

func main() {
	configPath := flag.String("config-path", "/run/secrets/config", "Please set in the config-path flag the path to config file")
	flag.Parse()

	fmt.Println("go-go-go", *configPath) //https://www.youtube.com/watch?v=rCJvW2xgnk0
}
