FROM golang:1.19 AS builder

ENV GO111MODULE on
ENV GOOS linux
ENV CGO_ENABLED 1
ENV GOARCH amd64

WORKDIR /usr/src

COPY . .

RUN go mod download && \
    go mod verify

RUN go build \
    -a \
    -tags netgo \
    -ldflags '-w -extldflags "-static"' \
    -o /usr/bin/server \
    -a cmd/server/main.go

FROM gcr.io/distroless/base-debian10:latest AS runner

COPY --from=builder /usr/bin/server /

CMD ["/server"]