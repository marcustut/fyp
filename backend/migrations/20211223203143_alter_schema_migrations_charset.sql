-- migrate:up
ALTER TABLE schema_migrations CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- migrate:down
ALTER TABLE schema_migrations CONVERT TO CHARACTER SET latin1 COLLATE latin1_swedish_ci;
