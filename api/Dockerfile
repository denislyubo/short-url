FROM golang:alpine as builder

COPY . /build

WORKDIR /build/api

RUN go mod tidy

RUN go build -o main .

#stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/api/.env /app/
COPY --from=builder /build/api/main /app/


WORKDIR /app

EXPOSE 3000

CMD ["./main"]