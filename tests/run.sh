#!/bin/sh

start() {
    pwd
    pkill -9 madrush
    ../bin/madrush &
    pid=$!
}

start &
curl -X GET localhost:9111/v1/subscribers
curl -X GET localhost:9111/v1/subscriber/abc
curl -X GET localhost:9112/health
curl -X GET localhost:9112/liveness
kill $pid
