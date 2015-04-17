package main

import (
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/rochacon/docker-cron/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

// Job is the definition of a CRON job
type Job struct {
	Id        string
	Interval  string
	Name      string
	Container docker.Config
}

// NewJob returns a Job instance parsed from a JSON reader
func NewJob(r io.Reader) (*Job, error) {
	j := &Job{}
	if err := json.NewDecoder(r).Decode(&j); err != nil {
		return j, err
	}
	return j, nil
}

// Run executes a Job
func (j *Job) Run() {
	started := time.Now()

	dcli, err := docker.NewClient(DOCKER_HOST)
	if err != nil {
		log.Printf("[%s] Error connecting to Docker instance: %q\n", j.Id, err)
		return
	}

	c, err := dcli.CreateContainer(docker.CreateContainerOptions{j.Name, &j.Container, nil})
	if err != nil {
		log.Printf("[%s] Error creating container: %q\n", j.Id, err)
		return
	}
	defer dcli.RemoveContainer(docker.RemoveContainerOptions{ID: c.ID})

	log.Printf("[%s] Starting job...\n", j.Id)
	if err := dcli.StartContainer(c.ID, &docker.HostConfig{}); err != nil {
		log.Printf("[%s] Error starting container: %q\n", j.Id, err)
		return
	}

	if _, err := dcli.WaitContainer(c.ID); err != nil {
		log.Printf("[%s] Error waiting container execution: %q\n", j.Id, err)
		return
	}

	log.Printf("[%s] Job finished in %s.\n", j.Id, time.Since(started))
}
