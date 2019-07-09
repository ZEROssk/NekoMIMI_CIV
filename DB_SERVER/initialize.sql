CREATE USER `api`@`%`;
GRANT INSERT,SELECT,DELETE ON `api_data_db`.* TO `api`@`%`;

CREATE DATABASE IF NOT EXISTS `api_data_db`;

CREATE TABLE IF NOT EXISTS `api_data_db`.`imgs` (
	`id`		INT			NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`img`		TEXT		NOT NULL,
	`add_date`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP
);

USE api_data_db;

INSERT INTO imgs (img) VALUES
('Twitter-test1'),
('Twitter-test2'),
('Twitter-test3'),
('Twitter-test4'),
('Twitter-test5'),
('Twitter-test6'),
('Twitter-test7'),
('Twitter-test8'),
('Twitter-test9'),
('Twitter-test10');

