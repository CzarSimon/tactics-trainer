-- +migrate Up
CREATE TABLE `cycle` (
    `id` VARCHAR(36) NOT NULL,
    `number` INTEGER NOT NULL,
    `problem_set_id` VARCHAR(36) NOT NULL,
    `current_puzzle_id` VARCHAR(36) NOT NULL,
    `completed_at` DATETIME,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`problem_set_id`) REFERENCES `problem_set` (`id`),
    FOREIGN KEY (`current_puzzle_id`) REFERENCES `puzzle` (`id`),
    UNIQUE(`problem_set_id`, `number`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- +migrate Down
DROP TABLE IF EXISTS `cycle`;