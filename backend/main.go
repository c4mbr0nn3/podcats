package main

import (
	"example/hello/config"
	"example/hello/db"
	"example/hello/jobs"
	"example/hello/server"
	"flag"
	"log"
	"os"
)

func main() {
	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		log.Println("Usage: podcats -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	log.Println("Environment: " + *env)
	config.Init(*env)
	db.Init()
	db.Migrate()
	jobs.InitCron()
	server.Init()
}
