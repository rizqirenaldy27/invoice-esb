FROM golang:1.20.5

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o invoice-esb .

CMD ["./invoice-esb"]