FROM golang:1.15.7-alpine3.13 as builder

LABEL maintainer="Quique <yourmail@mail.com>"

ENV GO111MODULE=on

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

## Distribution
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main ./

CMD [ "/app/main" ]
