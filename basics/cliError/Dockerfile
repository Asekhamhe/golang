
FROM golang:1.15-alpine

RUN mkdir /go/src/app
COPY ./errors.go /go/src/app
WORKDIR /go/src/app
RUN go build -o /go/src/app/errors errors.go
ENTRYPOINT [ "/go/src/app/errors" ]