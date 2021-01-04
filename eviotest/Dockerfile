FROM golang

WORKDIR /

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN wget http://www.linuxidc.com/files/repo/google-chrome.list -P /etc/apt/sources.list.d/
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub  |  apt-key add -

RUN apt-get update &&  apt-get install google-chrome-stable &&  apt-get install ttf-wqy-microhei ttf-wqy-zenhei xfonts-wqy
