FROM golang:latest

RUN apt-get update && \
	go get -u github.com/ant0ine/go-json-rest/rest

ADD ./ /go

CMD ["go", "run", "server.go"]
