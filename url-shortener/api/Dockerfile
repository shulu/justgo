FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go mod init url-shortener

RUN go mod tidy

RUN go build main.go

#stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app/

EXPOSE 3000

CMD ["./main"]
