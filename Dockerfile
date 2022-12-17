FROM golang:1.19.3 AS build 

RUN mkdir -p /go/src/whereami

ADD main.go /go/src/whereami
WORKDIR /go/src/whereami

RUN go mod init && go mod tidy && go build -o /whereami

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /whereami /whereami

EXPOSE 8080

ENTRYPOINT ["/whereami"]
