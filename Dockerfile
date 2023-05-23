FROM golang:1.20-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /periodictimestamps

CMD ["/periodictimestamps", "0.0.0.0", "8000"]
