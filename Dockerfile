FROM golang:latest

WORKDIR /go/src/task
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


