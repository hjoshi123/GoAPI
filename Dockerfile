FROM golang:latest as builder

LABEL maintainer="Hemant Joshi <joshi19981998@gmail.com>"

WORKDIR /go/src/github.com/hjoshi123/GORest

COPY . .
COPY .env .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/GORest .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash

WORKDIR /root/
RUN ls -a

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/GORest .
COPY --from=builder /go/src/github.com/hjoshi123/GORest/wait-for-it.sh .
RUN chmod +x wait-for-it.sh

EXPOSE 8080