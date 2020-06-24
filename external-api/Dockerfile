FROM golang:1.14-alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
ADD . ./
RUN go build -o api
EXPOSE 3001
CMD ["./api"]
