BEGIN;

ALTER TABLE categories
    DROP COLUMN is_active;

COMMIT;
