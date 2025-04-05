FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main ./cmd/main.go

FROM alpine:3
COPY --from=builder main /bin/main

EXPOSE 8060

ENTRYPOINT ["/bin/main"]