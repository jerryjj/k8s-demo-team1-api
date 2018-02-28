#### build stage ####
FROM golang:1.9 as golang

WORKDIR /build

ENV GOPATH=/build

ADD ./src ./src

RUN cd src/qvik.fi/api-service && go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/api-service qvik.fi/api-service

#### run stage ####
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=golang /build/bin/api-service /usr/local/bin/api-service

EXPOSE 8080

CMD ["/usr/local/bin/api-service"]
