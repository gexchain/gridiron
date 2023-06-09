FROM golang:1.19.5-bullseye AS build-env

WORKDIR /go/src/github.com/gridchain/gridiron

RUN apt-get update -y
RUN apt-get install git -y

COPY . .

RUN make build

FROM golang:1.19.5-bullseye

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/gridchain/gridiron/build/gridirond /usr/bin/gridirond

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["gridirond"]
