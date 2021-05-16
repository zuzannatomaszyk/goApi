FROM golang:1.14.6-alpine3.12 as builder
WORKDIR /goApi
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/goApi github.com/zuzannatomaszyk/goApi

FROM alpine
WORKDIR /serverFiles
COPY /db/migrations/** .
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /goApi/build/goApi /usr/bin/goApi
EXPOSE 8081 8081
ENTRYPOINT ["/usr/bin/goApi"]

