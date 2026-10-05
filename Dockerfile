FROM golang:1.26 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM debian:bookworm-slim
COPY --from=builder /src/resources/properties.yaml /resources/properties.yaml
COPY --from=builder /src/bin/ostinato /ostinato
ENTRYPOINT ["/ostinato"]
