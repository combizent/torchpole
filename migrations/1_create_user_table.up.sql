CREATE TABLE IF NOT EXISTS `user`
(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT,
    `username`  varchar(255) NOT NULL,
    `password`  varchar(255) NOT NULL,
    `nickname`  varchar(30)  NOT NULL,
    `email`     varchar(256) NOT NULL,
    `phone`     varchar(16)  NOT NULL,
    `createdAt` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb3;