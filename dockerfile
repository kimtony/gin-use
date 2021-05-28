FROM golang:1.15.5-alpine AS build
WORKDIR $GOPATH/src/gin-use
COPY . ./
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct 


RUN go mod init \
    && go mod tidy \
    && go build -ldflags="-s -w" -o /app/gin-use .

FROM alpine
    RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
    ENV TZ Asia/Shanghai \
        GIN_MODE=release 
    # 设置工作目录
    WORKDIR /data/app
    # 复制生成的可执行命令和一些配置文件
    COPY --from=build /app/gin-use .

EXPOSE 8081
CMD [ "./gin-use" ]