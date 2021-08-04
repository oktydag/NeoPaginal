FROM golang:1.16  as builder

RUN apk add --no-cache git
WORKDIR /app/neopaginal-command

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go builder -o ./out/neopaginal-command .

EXPOSE 5001
CMD ["./out/neopaginal-command"]