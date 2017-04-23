FROM golang:alpine

MAINTAINER ttpham0111

RUN apk add --update git

COPY . /go/src/ttpham0111

WORKDIR /go/src/ttpham0111/app
RUN go get -v && \
    go install -v

CMD ["app"]