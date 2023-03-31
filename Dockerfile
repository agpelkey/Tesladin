#syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY *.go ./

RUN go build -o /bin/tesladin 

EXPOSE 8080

CMD [ "/tesladin" ]