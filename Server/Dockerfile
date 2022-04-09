FROM golang:alpine as build

MAINTAINER yeol

ENV GOPROXY https://goproxy.cn,direct

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_server


FROM alpine:latest

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo 'Asia/Shanghai' >/etc/timezone

COPY --from=build /app/go_server /usr/bin/go_server
Run chmod +x /usr/bin/go_server

EXPOSE 25566

ENTRYPOINT ["go_server"]
