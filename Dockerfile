FROM golang:1.15-buster

RUN apt-get update
RUN apt-get install -y default-mysql-server default-mysql-client
RUN apt-get -y install vim
#RUN sed -i -e"s/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/" /etc/mysql/my.cnf

EXPOSE 80
EXPOSE 3306
COPY AvitoTest /go/src/AvitoTest
#COPY github.com /go/src/github.com
COPY ./start.sh /start.sh
RUN chmod +x /start.sh

CMD /start.sh
