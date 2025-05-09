FROM golang:1.24
WORKDIR /app
COPY . .
RUN go build -o server main.go
CMD ["./server"]