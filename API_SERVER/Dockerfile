FROM golang:latest

RUN apt-get update && \
	go get -u github.com/ant0ine/go-json-rest/rest && \
	go get github.com/go-sql-driver/mysql && \
	go get github.com/gorilla/mux

ADD ./ /go

RUN go build server.go

CMD ["./server"]
