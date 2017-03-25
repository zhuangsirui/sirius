FROM golang:latest
MAINTAINER zhuangsirui <zhuangsirui@gmail.com>

COPY . /go/src/sirius
RUN go install sirius
RUN mkdir /var/sirius
ENTRYPOINT ["/go/bin/sirius"]

EXPOSE 80
