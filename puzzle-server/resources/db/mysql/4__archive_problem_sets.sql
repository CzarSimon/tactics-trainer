-- +migrate Up
ALTER TABLE `problem_set`
ADD COLUMN `archived` BOOLEAN DEFAULT 'FALSE' NOT NULL
AFTER `user_id`;
-- +migrate Down
ALTER TABLE `problem_set` DROP COLUMN `archived`;