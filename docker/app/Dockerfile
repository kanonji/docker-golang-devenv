FROM golang:1.18-alpine
RUN apk add --no-cache git ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone
WORKDIR /go/src/app
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download
COPY ./src .
RUN go build -o /go/bin/app -v ./...
EXPOSE 8080
CMD go install github.com/cespare/reflex@latest && \
    reflex -r "\.go$" -s -- sh -c "cd /go/src/app && go build -o /go/bin/app -v ./... && /go/bin/app"
