FROM golang:1.11.1 as builder
WORKDIR /go/src/github.com/krezac/los-server/
COPY ./* ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/krezac/los-server/app .
EXPOSE 8080
CMD ["./app"]  
