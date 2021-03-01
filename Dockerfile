FROM bigrocs/golang-gcc:1.13 as builder
# 安装 odbc 依赖
RUN apk add --no-cache git make zip unixodbc unixodbc-dev freetds

ADD docker/etc_odbcinst.ini /etc/odbcinst.ini

WORKDIR /go/src/github.com/bxxinshiji/sql2000
COPY . .

ENV GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o bin/sql2000

FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 安装 odbc 依赖
RUN apk add --update unixodbc unixodbc-dev freetds

COPY --from=builder /go/src/github.com/bxxinshiji/sql2000/bin/sql2000 /usr/local/bin/
CMD ["sql2000"]
EXPOSE 8080