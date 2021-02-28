FROM bigrocs/golang-gcc:1.13 as builder

WORKDIR /go/src/github.com/lecex/sql2000
COPY . .

ENV GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o bin/sql2000

FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/github.com/lecex/sql2000/bin/sql2000 /usr/local/bin/
CMD ["sql2000"]
EXPOSE 8080