FROM golang:1.17.7 AS builder
WORKDIR /go/src/github.com/ddld93/abedmis/
RUN go get -d -v golang.org/x/net/html  
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/ddld93/abedmis/main ./
CMD ["./main"] 
 