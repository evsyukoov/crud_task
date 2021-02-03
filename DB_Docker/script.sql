CREATE DATABASE IF NOT EXISTS avito_test;
use avito_test;

CREATE TABLE avito_test.statistic (date DATE NOT NULL, views INT NOT NULL, clicks INT NOT NULL, cost DECIMAL(8,2) NOT NULL, 
PRIMARY KEY (date));


-- GRANT ALL PRIVILEGES ON avito_test.* TO 'admin'@'localhost' WITH GRANT OPTION
-- echo "update mysql.user set host='%' where host='localhost'" | mysql -u root --skip-password
-- echo "UPDATE mysql.user SET Password=PASSWORD('1111') WHERE User='root' AND Host='%'" |  mysql -u root --skip-password
