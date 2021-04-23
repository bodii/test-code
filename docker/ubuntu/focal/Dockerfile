FROM ubuntu:20.04
MAINTAINER bodii <bodii@example.com>
ENV REFRESHED_AT 2021-04-23


# setting aliyun source
ADD sources_focal.list /etc/apt/sources.list

# setting timezone
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update \
    && apt-get install -yqq tzdata \
    && ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
