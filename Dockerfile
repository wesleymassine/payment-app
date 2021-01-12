FROM golang:latest

LABEL maintainer="Wesley Santos - wesley.massine@gmail.com | github.com/wesleymassine"

WORKDIR /api

COPY . .

RUN go mod download

RUN go build

CMD ["./payment_app"]
