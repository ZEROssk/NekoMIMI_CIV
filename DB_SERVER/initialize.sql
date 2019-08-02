CREATE USER `api`@`%`;
GRANT INSERT,SELECT,DELETE ON `api_data_db`.* TO `api`@`%`;

CREATE DATABASE IF NOT EXISTS `api_data_db`;

CREATE TABLE IF NOT EXISTS `api_data_db`.`twimg_data` (
	`ID`		INT			NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`TwiID`		VARCHAR(15)	NOT NULL,
	`FileName`	VARCHAR(62)	NOT NULL,
	`CreatedAt`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP
);

USE api_data_db;

INSERT INTO twimg_data (TwiID, FileName) VALUES
('aaa','test-01'),
('aaa','test-02'),
('abb','test-03'),
('abb','test-04'),
('abb','test-05'),
('abb','test-06'),
('abb','test-07'),
('aaa','test-08'),
('aaa','test-09'),
('abc','test-10'),
('abc','test-11'),
('abc','test-12'),
('abc','test-13'),
('abc','test-14'),
('abc','test-15'),
('abc','test-16'),
('abc','test-17'),
('aaa','test-18'),
('aaa','test-19'),
('aaa','test-20'),
('aaa','test-21'),
('aaa','test-22'),
('aaa','test-23'),
('aaa','test-24'),
('aaa','test-25'),
('aaa','test-26'),
('aaa','test-27'),
('aaa','test-28'),
('abc','test-29'),
('abc','test-30'),
('abc','test-31'),
('abc','test-32'),
('abc','test-33'),
('abb','test-34'),
('abb','test-35'),
('abb','test-36'),
('abb','test-37'),
('abb','test-38'),
('abb','test-39'),
('abc','test-40'),
('abc','test-41'),
('abc','test-42'),
('abc','test-43'),
('abc','test-44'),
('abc','test-45'),
('abc','test-46'),
('abc','test-47'),
('abb','test-48'),
('abb','test-49'),
('abb','test-50'),
('abb','test-51'),
('abb','test-52'),
('abb','test-53'),
('abc','test-54'),
('aaa','test-55'),
('aaa','test-56'),
('aaa','test-57'),
('abc','test-58');

