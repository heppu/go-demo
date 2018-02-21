FROM golang:1.10 AS builder
WORKDIR /go/src/github.com/heppu/go-demo/
COPY . ./
RUN make build

FROM scratch
COPY --from=builder /go/src/github.com/heppu/go-demo/app .
CMD ["./app"]
