FROM golang:1.16-buster as builder

# ENV GOOS linux
# ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

# RUN go build -o app

# FROM alpine:3.14 as production

FROM gcr.io/distroless/base-debian10

# RUN apk add --no-cache ca-certificates

WORKDIR /

COPY --from=builder /docker-gs-ping /docker-gs-ping

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]