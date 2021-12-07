FROM golang:latest

ENV GOPATH=

ADD ./ /go

RUN apt-get update && \
	export GO111MODULE=on && \
	go mod init main && \
	go mod tidy && \
	go build server.go

CMD ["./server"]
