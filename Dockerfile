FROM archlinux
LABEL maintainer Yaks Souza <pl4nk1ton@gmail.com>
RUN pacman -Sy go --noconfirm

RUN mkdir /api
ADD . /api

EXPOSE 8000

WORKDIR /api

CMD go run . >> out.log 2>> out.log
