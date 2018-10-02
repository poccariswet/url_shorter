FROM golang:latest

WORKDIR /go/src/url_shorter
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["url_shorter"]
