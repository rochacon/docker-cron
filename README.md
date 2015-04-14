# Docker CRON

`docker-cron` is a Docker base image that schedules containers in a time interval bases.


# Building your cron image

Create a `Dockerfile` with the following contents:


```
FROM rochacon/docker-cron
```


Now, create a folder `jobs` and place your jobs definitions in it. Here is a sample job definition:


```json
// This is a simple Hello world job that will run every 10 seconds
{
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
}
```


The `Interval` field must be the cron expression. Reference: http://godoc.org/github.com/robfig/cron#hdr-CRON_Expression_Format
The `Name` field is optional and can be used to guarantee that only one instance of the job will run at a time.
The `Container` field is the description of the container and may include these fields: http://godoc.org/github.com/fsouza/go-dockerclient#Config

Now build your cron image:


```
docker build -t my-docker-cron .
```


# Running

Now that we built our cron image, let's run it. `docker-cron` uses the `DOCKER_HOST` environment variable to comunicate with your Docker instance and defaults to `unix://var/run/docker.sock`.

To run the `docker-cron` with a local Docker instance:


```
docker run -d --name docker-cron -v /var/run/docker.sock:/var/run/docker.sock:rw my-docker-cron
```


Or with a remote Docker instance:


```
docker run -d --name docker-cron -e DOCKER_HOST=tcp://remote-docker.domain:2375 my-docker-cron
```


:warning: Currently TLS connection is not supported.



# License

MIT
