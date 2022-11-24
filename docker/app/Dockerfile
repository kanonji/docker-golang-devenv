FROM golang:1.19-alpine
RUN apk add --no-cache git ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone
WORKDIR /go/src/app
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download
COPY ./src .
EXPOSE 8080
CMD go install github.com/cespare/reflex@latest && \
    reflex -r "(\.go$|go\.mod$)" -s go run /go/src/app/cmd/main.go
