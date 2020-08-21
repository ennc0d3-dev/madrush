FROM golang:1.15 AS build
WORKDIR /go/src
COPY cmd/ internal/ vendor/ go.mod go.sum  ./

RUN go mod download && \
 go install  -v ./...

FROM scratch AS runtime
COPY --from=build /go/bin/madrush ./madrush
EXPOSE 8080/tcp
ENTRYPOINT ["./madrush"]
