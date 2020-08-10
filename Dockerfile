# Build:
#    docker build -t cosasdepuma/elliot .
# Run:
#    sudo docker run -d -p 8008:8008 --name elliot cosasdepuma/elliot

FROM golang:alpine AS build

WORKDIR /go/src/app
COPY . /go/src/app

RUN sed -i 's/^const Host = .*$/const Host = "0.0.0.0"/g' /go/src/app/pkg/config/config.go && \
    sed -i 's/^const Port = .*$/const Port = 8008/g' /go/src/app/pkg/config/config.go && \
    sed -i 's/^const DBPass = .*$/const DBPass = "d0cK3r"/g' /go/src/app/pkg/config/config.go

RUN apk update && \
    apk add --no-cache --upgrade --virtual .requirements build-base && \
    CGO_ENABLED=1 go build -v -o elliot main.go && \
    apk del .requirements

# ================================

FROM alpine

ENV REDISPASS foobar

RUN apk update && \
    apk add redis && \
    sed -i 's/^# requirepass .*$/requirepass d0ck3r/g' /etc/redis.conf

RUN addgroup -S elliot && adduser -s /sbin/nologin -S elliot -G elliot
USER elliot

WORKDIR /app
COPY --from=build /go/src/app/elliot /bin/elliot
COPY --from=build /go/src/app/frontend /app/frontend

EXPOSE 8008

CMD ["/bin/elliot"]