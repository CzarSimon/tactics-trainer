-- +migrate Up
ALTER TABLE `problem_set`
ADD COLUMN `archived` BOOLEAN NOT NULL DEFAULT 0
AFTER `user_id`;
-- +migrate Down
ALTER TABLE `problem_set` DROP COLUMN `archived`;