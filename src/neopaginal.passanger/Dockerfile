FROM golang:1.16  as builder

RUN apk add --no-cache git
WORKDIR /app/neopaginal-passanger

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go builder -o ./out/neopaginal-passanger .

EXPOSE 5002
CMD ["./out/neopaginal-passanger"]