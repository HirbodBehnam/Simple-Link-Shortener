CREATE DATABASE `link_shortener`;
CREATE TABLE `link_shortener`.`links`
(
    `key`  VARCHAR(16) NOT NULL COMMENT 'The key to url',
    `link` TEXT        NOT NULL COMMENT 'The link to address',
    PRIMARY KEY (`key`)
) ENGINE = InnoDB;