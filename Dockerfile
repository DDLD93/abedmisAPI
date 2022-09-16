FROM golang:1.17.7 AS builder
WORKDIR /go/src/github.com/ddld93/abedmis/
RUN go get -d -v golang.org/x/net/html  
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
COPY app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/ddld93/abedmis/app ./
CMD ["./app"] 