# syntax=docker/dockerfile:1

FROM golang:1.19-buster AS build

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

ADD . /go/src/app/
COPY *.go ./

RUN go build -o /url-shorten

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY *.sql ./
COPY config.sample.json ./config.json

COPY --from=build /url-shorten /url-shorten

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/url-shorten"]
CMD [ "config.json" ]