-- Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file. The original repo for
-- this file is https://github.com/rppkg/torchpole.

CREATE DATABASE IF NOT EXISTS `torchpole`;

USE `torchpole`;

CREATE TABLE IF NOT EXISTS `user`
(
    `id`        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username`  varchar(255) NOT NULL,
    `password`  varchar(255) NOT NULL,
    `nickname`  varchar(30)  NOT NULL,
    `email`     varchar(256) NOT NULL,
    `phone`     varchar(16)  NOT NULL,
    `createdAt` timestamp    NOT NULL DEFAULT current_timestamp(),
    `updatedAt` timestamp    NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp (),
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;