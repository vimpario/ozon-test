FROM golang:latest as builder
WORKDIR /go/src/ozon-test
COPY go.mod go.sum /go/src/ozon-test/
RUN go mod download
COPY . /go/src/ozon-test
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o build/ozon-test ./cmd/main.go

FROM alpine
RUN  apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/ozon-test/build/ozon-test /usr/bin/ozon-test
COPY --from=builder /go/src/ozon-test/.env .
EXPOSE 9000 9000
ENTRYPOINT ["/usr/bin/ozon-test"]