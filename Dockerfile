FROM golang:1.18-alpine3.15 as builder
LABEL maintainer="Chaitanya Vemprala <cvemprala@gmail.com>"

WORKDIR /artifacts/
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o books .

FROM scratch
EXPOSE 8080

WORKDIR /app/
COPY --from=builder /artifacts/books .

ENTRYPOINT ["./books"]