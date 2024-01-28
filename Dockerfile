FROM golang:1.21 AS builder
# 用于go mod tidy
ENV GOPROXY https://goproxy.cn,direct
# 时区
ENV TZ Asia/Shanghai
# 禁用cgo，目前不理解作用
ENV CGO_ENABLED 0
# 创建并指定工作目录
RUN mkdir -p /app
WORKDIR /app

ADD . /app
RUN go mod tidy
RUN make build


FROM alpine
# Ca证书和时区
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ Asia/Shanghai
ENV service api

WORKDIR /app
# 原路径到目标路径
COPY --from=builder /app/cmd /app/cmd
COPY --from=builder /app/config /app/config

CMD ["sh", "-c", "cd cmd/${service}/output && sh bootstrap.sh"]
#docker run -d --name "ice-pomelo-api" -e service="api" --net=host ice-pomelo
#docker run -d --name "ice-pomelo-user" -e service="user" --net=host ice-pomelo
#docker run -d --name "ice-pomelo-tiny_id" -e service="tiny_id" --net=host ice-pomelo