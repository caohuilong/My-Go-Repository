CREATE DATABASE IF NOT EXISTS `staff`;
CREATE USER chl identified by 'chl123';
GRANT select, insert, update, create, delete on staff.* to 'chl'@'%' with grant option;
FLUSH privileges;
USE staff;
CREATE TABLE `ecf_staff`(
	id INT UNSIGNED AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	position  VARCHAR(50)NOT NULL,
	phonenumber VARCHAR(11) NOT NULL,
	PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
