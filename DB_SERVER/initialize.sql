CREATE USER `api`@`%`;
GRANT INSERT,SELECT,DELETE ON `api_data_db`.* TO `api`@`%`;

CREATE DATABASE IF NOT EXISTS `api_data_db`;

CREATE TABLE IF NOT EXISTS `api_data_db`.`twi_data` (
	`id`		INT			NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`twi_id`	TEXT		NOT NULL,
	`file_name`	TEXT		NOT NULL,
	`add_date`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP
);

USE api_data_db;

INSERT INTO twi_data (twi_id, file_name) VALUES
('a','Twitter-test1'),
('a','Twitter-test2'),
('b','Twitter-test3'),
('b','Twitter-test4'),
('b','Twitter-test5'),
('c','Twitter-test6'),
('d','Twitter-test7'),
('ee','Twitter-test8'),
('ee','Twitter-test9'),
('abc','Twitter-test10');

