#FROM golang:1.11
#
## Add Maintainer Info
#LABEL maintainer="Sjoerd During"
#
## Set the Current Working Directory inside the container
#WORKDIR $GOPATH/src/github.com/SDur/ops-planner
#
## Copy everything from the current directory to the PWD(Present Working Directory) inside the container
#COPY . .
#
#RUN go get -d -v ./...
#
## Install the package
#RUN go install -v ./...
#
## This container exposes port 8080 to the outside world
#EXPOSE 8080
#
## Run the executable
#CMD ["ops-planner"]
FROM golang:onbuild
EXPOSE 8080