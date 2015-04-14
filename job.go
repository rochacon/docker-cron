package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/rochacon/docker-cron/Godeps/_workspace/src/github.com/fsouza/go-dockerclient"
)

// Job is the definition of a CRON job
type Job struct {
	Id        string
	Interval  string
	Name      string
	Container docker.Config
}

// FromFile returns a Job instance parsed from a JSON file
func FromFile(filename string) (*Job, error) {
	j := &Job{}

	fp, err := os.Open(filename)
	if err != nil {
		return j, err
	}

	if err := json.NewDecoder(fp).Decode(&j); err != nil {
		return j, err
	}

	return j, nil
}

// Run executes a Job
func (j *Job) Run() {
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
	log.Printf("[%s] Job finished.\n", j.Id)
}
