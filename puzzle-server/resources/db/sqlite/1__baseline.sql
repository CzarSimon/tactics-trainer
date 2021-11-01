-- +migrate Up
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
CREATE TABLE `problem_set` (
    `id` VARCHAR(36) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `themes` VARCHAR(255) NOT NULL,
    `rating_interval` VARCHAR(255) NOT NULL,
    `user_id` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `problem_set_puzzle` (
    `id` VARCHAR(36) NOT NULL,
    `puzzle_id` VARCHAR(36) NOT NULL,
    `problem_set_id` VARCHAR(36) NOT NULL,
    `number` INTEGER NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`puzzle_id`) REFERENCES `puzzle` (`id`),
    FOREIGN KEY (`problem_set_id`) REFERENCES `problem_set` (`id`)
);
-- +migrate Down
DROP TABLE IF EXISTS `problem_set_puzzle`;
DROP TABLE IF EXISTS `problem_set`;
DROP TABLE IF EXISTS `puzzle`;