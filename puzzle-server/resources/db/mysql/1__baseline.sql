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
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- +migrate Down
DROP TABLE IF EXISTS `puzzle`;
