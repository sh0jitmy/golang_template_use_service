FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN ls -l -d $(find `pwd`) 
RUN go build -o app ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
