FROM golang:1.17-alpine AS build
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk --no-cache add tzdata ca-certificates && update-ca-certificates

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
    #GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

WORKDIR /src/
COPY ./server/go.mod  ./
COPY ./server/go.sum ./
RUN go mod download

COPY ./server/ /src/
RUN CGO_ENABLED=0 go build -o /bin/plant-dashboard


FROM alpine as final

WORKDIR /app/

COPY --from=build /bin/plant-dashboard /app/plant-dashboard
COPY --from=build /src/resource ./resource/
COPY --from=build /src/config.yaml ./

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


ENV TZ=Asia/Shanghai

EXPOSE 8888
ENTRYPOINT ["/app/plant-dashboard", "-c", "/app/config.yaml"]