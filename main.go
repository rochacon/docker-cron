package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/rochacon/docker-cron/Godeps/_workspace/src/github.com/robfig/cron"
)

// DOCKER_HOST points to the Docker daemon socket
var DOCKER_HOST string

func main() {
	jobsDir := flag.String("jobs", "jobs", "Directory containing job definitions.")
	flag.Parse()

	// Get Docker host from environment
	DOCKER_HOST = os.Getenv("DOCKER_HOST")
	if DOCKER_HOST == "" {
		DOCKER_HOST = "unix:///var/run/docker.sock"
	}

	// Get jobs descriptions filenames
	jobs, err := filepath.Glob(fmt.Sprintf("%s/*", *jobsDir))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Setup cron
	cr := cron.New()
	for _, filename := range jobs {
		j, err := FromFile(filename)
		if err != nil {
			fmt.Printf("Error parsing job from %q: %q\n", filename, err)
			os.Exit(2)
		}
		fmt.Printf("Registering job %q for %q\n", j.Id, j.Interval)
		err = cr.AddJob(j.Interval, j)
		if err != nil {
			fmt.Printf("Error registering job %q: %q\n", j.Id, err)
			os.Exit(2)
		}
	}
	cr.Start()

	// Wait SIGINT or SIGKILL to quit app
	fmt.Println("Started, press CTRL+C or send a SIGKILL to quit.")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)
	<-exit
	cr.Stop()
	// TODO clean running jobs.
}
