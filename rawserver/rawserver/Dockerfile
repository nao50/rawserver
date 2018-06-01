FROM golang:latest
WORKDIR ${GOPATH}/
RUN apt-get update -y && apt-get install -y vim tcpdump iperf3 traceroute strace

ADD . ${GOPATH}