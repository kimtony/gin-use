FROM golang:1.15.5-alpine AS build-env
WORKDIR $GOPATH/src/gin-use
COPY . ./
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct 
RUN go mod init \
    && go mod tidy \
    && go build -o gin-use .

EXPOSE 8080
CMD [ "./gin-use" ]