FROM golang:1.11

USER nobody

RUN mkdir -p /go/src/github.com/SDur/ops-planner
WORKDIR /go/src/github.com/SDur/ops-planner

COPY . .
RUN go get -d -v ./...
RUN go build

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["ops-planner"]
