FROM golang:1.19

WORKDIR /app

COPY main.go .

RUN go build -o server main.go

EXPOSE 3000

CMD ["./server"]