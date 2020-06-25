FROM golang:1.14-alpine as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
ADD . ./
RUN go build -o app
EXPOSE 3000
CMD ["./app"]
