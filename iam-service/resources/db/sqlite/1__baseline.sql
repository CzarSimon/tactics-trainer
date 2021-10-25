-- +migrate Up
CREATE TABLE `key_state` (
    `id` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `key_encryption_key` (
    `id` INTEGER NOT NULL,
    `state` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`state`) REFERENCES `key_state` (`id`)
);
CREATE TABLE `role` (
    `name` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`name`)
);
CREATE TABLE `user_account` (
    `id` VARCHAR(50) NOT NULL,
    `username` VARCHAR(50) NOT NULL,
    `role` VARCHAR(50) NOT NULL,
    `password` VARCHAR(256) NOT NULL,
    `salt` VARCHAR(64) NOT NULL,
    `data_encryption_key` VARCHAR(255) NOT NULL,
    `key_encryption_key_id` INTEGER NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`username`),
    UNIQUE(`salt`),
    UNIQUE(`data_encryption_key`),
    FOREIGN KEY (`role`) REFERENCES `role` (`name`),
    FOREIGN KEY (`key_encryption_key_id`) REFERENCES `key_encryption_key` (`id`)
);
CREATE TABLE `puzzle` (
    `id` VARCHAR(36) NOT NULL,
    `external_id` VARCHAR(10) NOT NULL,
    `fen` VARCHAR(255) NOT NULL,
    `moves` VARCHAR(255) NOT NULL,
    `rating` INTEGER NOT NULL,
    `rating_deviation` INTEGER NOT NULL,
    `popularity` INTEGER NOT NULL,
    `themes` VARCHAR(255) NOT NULL,
    `game_url` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);
INSERT INTO `role`(`name`, `created_at`)
VALUES ('USER', CURRENT_TIMESTAMP);
INSERT INTO `key_state`(`id`, `created_at`)
VALUES ('ACTIVE', CURRENT_TIMESTAMP),
    ('NEXT', CURRENT_TIMESTAMP),
    ('DEACTIVATED', CURRENT_TIMESTAMP);
INSERT INTO `key_encryption_key`(`id`, `state`, `created_at`)
VALUES (0, 'ACTIVE', CURRENT_TIMESTAMP);
-- +migrate Down
DROP TABLE IF EXISTS `user_account`;
DROP TABLE IF EXISTS `role`;
DROP TABLE IF EXISTS `key_encryption_key`;
DROP TABLE IF EXISTS `key_state`;