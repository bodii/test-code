FROM ubuntu:20.10
MAINTAINER bodii "bodii@example.com"
ENV REFRESHED_AT 2021-04-23


# setting aliyun source
ADD sources_groovy.list /etc/apt/sources.list

# setting timezone
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update 
RUN apt-get install --assume-yes apt-utils 
RUN apt-get install -yqq tzdata \
    && ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && dpkg-reconfigure --frontend noninteractive tzdata \
    && rm -rf /var/lib/apt/lists/*
