FROM golang:1.26
WORKDIR /app
COPY . .
RUN make build
ENTRYPOINT [ "/app/bin/ostinato" ]
