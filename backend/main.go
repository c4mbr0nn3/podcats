package main

import (
	"example/hello/config"
	"example/hello/server"
	"flag"
	"fmt"
	"os"
)

func main() {
	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Environment: " + *env)
	config.Init(*env)
	//db.Init()
	server.Init()
}
