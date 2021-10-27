# stage I - khusus build dengan envinroment yang sama
FROM golang:1.17.2-alpine3.14 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main
# EXPOSE 8080
# CMD ["/app/main"]

# stage 2
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/app.config.json .
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]