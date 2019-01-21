FROM golang:1.11-alpine AS build-env

# To copy complete directory
COPY ./ /go/src/github.com/ShyftNetwork/go-empyrean
WORKDIR /go/src/github.com/ShyftNetwork/go-empyrean

RUN \
  apk add --update git make gcc musl-dev linux-headers ca-certificates && \
  (cd /go/src/github.com/ShyftNetwork/go-empyrean && make geth && make bootnode) && \
  cp -v /go/src/github.com/ShyftNetwork/go-empyrean/build/bin/geth /bin && \
  cp -v /go/src/github.com/ShyftNetwork/go-empyrean/build/bin/bootnode /bin

FROM alpine:3.8
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /go-empyrean/
COPY --from=build-env /bin/geth /bin/
COPY --from=build-env /bin/bootnode /bin/

EXPOSE 8545 8546 31333 31333/udp 8081
CMD ["geth"]
