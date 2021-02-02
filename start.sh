
#!/bin/sh
service mysql start

echo "CREATE DATABASE avito_test;" | mysql -u root --skip-password
echo "GRANT ALL PRIVILEGES ON avito_test.* TO 'root'@'localhost' WITH GRANT OPTION;" | mysql -u root --skip-password
echo "UPDATE mysql.user SET Password=PASSWORD('1111') WHERE User='root' AND Host='localhost'" |  mysql -u root --skip-password

 echo "CREATE TABLE avito_test.statistic (date DATE NOT NULL, \
 views INT NOT NULL, \
 сlicks INT NOT NULL, \
 cost DECIMAL(8,2) NOT NULL, \
 PRIMARY KEY (date));" | mysql -u root --skip-password

cd /go/src/AvitoTest
make run
sh