DROP DATABASE IF EXISTS los;
CREATE DATABASE IF NOT EXISTS los;
USE los;

-- not compatible DROP USER IF EXISTS los@localhost;
CREATE USER 'los'@localhost IDENTIFIED BY 'los';
GRANT USAGE ON los.* TO 'los'@localhost IDENTIFIED BY PASSWORD 'los';
FLUSH PRIVILEGES;

SHOW GRANTS FOR 'los'@localhost;   