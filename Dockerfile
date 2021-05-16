FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/zuzannatomaszyk/goApi/
WORKDIR /go/src/github.com/zuzannatomaszyk/goApi
RUN go mod download
COPY . /go/src/github.com/zuzannatomaszyk/goApi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/goApi github.com/zuzannatomaszyk/goApi

FROM alpine
WORKDIR /serverFiles
COPY /db/migrations/** .
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/zuzannatomaszyk/goApi/build/goApi /usr/bin/goApi
EXPOSE 8081 8081
ENTRYPOINT ["/usr/bin/goApi"]

