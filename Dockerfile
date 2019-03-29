FROM golang:1.12

USER nobody

RUN mkdir -p /go/src/github.com/SDur/ops-planner
WORKDIR /go/src/github.com/SDur/ops-planner

COPY . /go/src/github.com/SDur/ops-planner
RUN go mod init
RUN go build

CMD ["./golang-ex"]
