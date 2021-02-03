use avito_test;

CREATE TABLE IF Not Exists avito_test.statistic (date DATE NOT NULL, views INT NOT NULL, clicks INT NOT NULL, cost DECIMAL(8,2) NOT NULL,
PRIMARY KEY (date));
GRANT ALL PRIVILEGES ON avito_test.* TO 'admin'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
