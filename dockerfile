# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /loiter

EXPOSE 9500
EXPOSE 9501

CMD [ "/loiter" ]