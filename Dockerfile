FROM golang:latest

ENV GOBIN /go/bin

WORKDIR /go/src/todo
COPY . .

RUN go get -d -v ./...
RUN go install ./...

CMD ["todo"]
