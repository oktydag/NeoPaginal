FROM golang:1.16  as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


FROM alpine:3.12
COPY --from=builder /app/. /app/.

WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["./neopaginal-command"]
