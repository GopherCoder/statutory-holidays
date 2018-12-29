FROM ubuntu:latest
MAINTAINER "115143589@qq.com"


WORKDIR /home

COPY holiday /home/


CMD ["bash", "-c", "/home/holiday"]

