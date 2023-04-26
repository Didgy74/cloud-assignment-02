FROM golang:1.19

LABEL maintainer="mail@domain.tld"

COPY ./cmd /go/src/app/cmd
COPY ./handlers /go/src/app/handlers
COPY ./utils /go/src/app/utils
COPY ./go.mod /go/src/app/go.mod

WORKDIR /go/src/app/cmd/server

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server

EXPOSE 8080

CMD ["./server"]