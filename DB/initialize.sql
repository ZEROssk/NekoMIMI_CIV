CREATE USER `imgs_data`@`%`;
GRANT INSERT,SELECT,UPDATE,DELETE ON `imgs_data_db`.* TO `imgs_data`@`%`;

CREATE DATABASE IF NOT EXISTS `imgs_data_db`;

CREATE TABLE IF NOT EXISTS `imgs_data_db`.`imgs` (
	`id`		INT			NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`img`		TEXT		NOT NULL,
	`add_date`	TIMESTAMP	NOT NULL DEFAULT CURRENT_TIMESTAMP
);

USE imgs_data_db;

INSERT INTO imgs (img) VALUES
('Twitter-1141681496014520323-mzkzmm-D9gRRRHVUAAWTuN.jpg'),
('Twitter-998662632864272384-satokazu3232-Ddv1ZxdU0AE58Uz.jpg'),
('Twitter-998188062239744000-fuku_fox-DdpAd6xVAAA-wqq.jpg'),
('Twitter-998129394336940032-satokazu3232-DdoRKE3VMAAJ2zi.jpg'),
('Twitter-992420865022218240-fuku_fox-DcW10O7V4AAsN_l.jpg'),
('Twitter-982551120588324864-moriguruta-DaK4xqvUQAA_IG5.png'),
('Twitter-981127631348293633-hakuishiaoi-DZ2qWcaU8AE6p0s.jpg'),
('Twitter-969257592152190976-kuiuji_tuki-DXN-gezU8AAOaMq.jpg'),
('Twitter-942001317194510336-kagarifire_-DRKnza_VwAAXN6c.png'),
('Twitter-911944312098885632-N_A_Z_-DKfgcVcVAAEWSEM.png');

