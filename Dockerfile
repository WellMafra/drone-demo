FROM golang:1.9
RUN mkdir -p /go/src/github.com/WellMafra/drone-demo
COPY /go/src/github.com/WellMafra/drone-demo /go/src/github.com/WellMafra/drone-demo
WORKDIR /go/src/github.com/WellMafra/drone-demo
ENV GOPATH /go