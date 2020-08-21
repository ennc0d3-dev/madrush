#!/bin/bash

start() {
    pkill -9 madrush
    ../bin/madrush &
    sleep 2
}

start
curl -X GET localhost:9111/v1/subscribers
curl -X GET localhost:9111/v1/subscriber/abc
curl -X GET localhost:9112/health
curl -X GET localhost:9112/liveness
pkill -9 madrush
