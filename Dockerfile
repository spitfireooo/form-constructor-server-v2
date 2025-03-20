FROM golang:1.23

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o server-v2 ./cmd/main.go

CMD ["./server-v2"]