FROM golang:1.4-onbuild
ONBUILD COPY jobs/ /docker-jobs
ENTRYPOINT ["/go/bin/app"]
CMD ["-jobs", "/docker-jobs"]
