-- +migrate Up
ALTER TABLE `problem_set`
ADD `archived` BOOLEAN DEFAULT 'FALSE' NOT NULL;
-- +migrate Down
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;
ALTER TABLE `problem_set`
    RENAME TO `_problem_set_backup`;
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
INSERT INTO `problem_set`(
        `id`,
        `name`,
        `description`,
        `themes`,
        `rating_interval`,
        `user_id`,
        `created_at`,
        `updated_at`
    )
SELECT `id`,
    `name`,
    `description`,
    `themes`,
    `rating_interval`,
    `user_id`,
    `created_at`,
    `updated_at`
FROM `_problem_set_backup`;
COMMIT;
PRAGMA foreign_keys = on;