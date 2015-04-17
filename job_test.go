package main

import (
	"strings"
	"testing"
)

const jobDescription = `{
	"Id": "docker-hello-world-cron-id",
	"Interval": "*/10 * * * * *",
	"Name": "docker-hello-world",
	"Container": {
		"Cmd": ["echo", "$DISCLAIMER"],
		"Image": "busybox",
		"Env": [
			"DISCLAIMER=Hello World"
		]
	}
}`

func TestNewJob(t *testing.T) {
	description := strings.NewReader(jobDescription)
	job, err := NewJob(description)
	if err != nil {
		t.Errorf(err.Error())
	}
	if job.Id != "docker-hello-world-cron-id" {
		t.Errorf(`job.Id != "docker-hello-world-cron-id"`)
	}
}

func TestNewJobBadReaderContent(t *testing.T) {
	description := strings.NewReader("invalid")
	_, err := NewJob(description)
	if err == nil {
		t.Errorf("No error for invalid JSON.")
	}
}
