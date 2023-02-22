FROM golang:alpine3.17 as builder

ENV WORKSPACE=${GOPATH}/src/app
WORKDIR ${WORKSPACE}

COPY . .

RUN apk add --no-cache \
    make

RUN make build-linux

# cleaning image
FROM golang:alpine3.17

ENV WORKSPACE=${GOPATH}/src/app
WORKDIR ${WORKSPACE}

COPY --from=builder ${WORKSPACE}/myapi ${WORKSPACE}/myapi

CMD ["./myapi"]
