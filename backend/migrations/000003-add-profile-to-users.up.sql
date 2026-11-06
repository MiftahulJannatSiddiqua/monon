-- +migrate Up
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS bio TEXT,
    ADD COLUMN IF NOT EXISTS city VARCHAR(100),
    ADD COLUMN IF NOT EXISTS personality_type VARCHAR(20)
        CHECK (personality_type IN ('introvert', 'extrovert'));