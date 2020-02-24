#!/bin/sh

start() {
    pwd
    ../bin/madrush &
    pid=$!
}

start &
curl -X GET localhost:8080/v1/subscribers
curl -X GET localhost:8080/v1/subscriber/abc
kill $pid
