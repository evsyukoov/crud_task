FROM golang:1.15-buster

RUN apt-get update
RUN apt-get -y install vim

EXPOSE 80
COPY ./Application /go/src/Avito/Application

COPY ./start.sh /start.sh
RUN chmod +x /start.sh

CMD /start.sh