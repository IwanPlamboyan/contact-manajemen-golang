DROP INDEX IF EXISTS idx_refresh_tokens_active;
DROP INDEX IF EXISTS idx_refresh_tokens_expires_at;
DROP INDEX IF EXISTS idx_refresh_tokens_user_id;

ALTER TABLE refresh_tokens
DROP CONSTRAINT IF EXISTS uq_refresh_tokens_token_hash;

ALTER TABLE refresh_tokens
DROP CONSTRAINT IF EXISTS fk_refresh_tokens_user;

DROP TABLE IF EXISTS refresh_tokens;
