# syntax=docker/dockerfile:1

FROM golang:1.15-alpine
WORKDIR /app
COPY app/main.go ./
RUN go build -o /bitcoin-price
EXPOSE 8080
CMD [ "/bitcoin-price" ]