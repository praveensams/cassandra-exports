FROM golang:1.16-alpine

WORKDIR /mnt

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /docker-gs-ping

EXPOSE 10000

CMD [ "/docker-gs-ping" ]

