-- +migrate Up
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
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `problem_set_puzzle` (
    `id` VARCHAR(36) NOT NULL,
    `puzzle_id` VARCHAR(36) NOT NULL,
    `problem_set_id` VARCHAR(36) NOT NULL,
    `number` INTEGER NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`puzzle_id`) REFERENCES `puzzle` (`id`),
    FOREIGN KEY (`problem_set_id`) REFERENCES `problem_set` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE INDEX `puzzle_rating_idx` ON `puzzle` (`rating`);
CREATE INDEX `puzzle_popularity_idx` ON `puzzle` (`popularity`);
-- +migrate Down
ALTER TABLE `puzzle` DROP INDEX `puzzle_popularity_idx`;
ALTER TABLE `puzzle` DROP INDEX `puzzle_rating_idx`;
DROP TABLE IF EXISTS `problem_set_puzzle`;
DROP TABLE IF EXISTS `problem_set`;