-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY REFERENCES users(guid) ON DELETE CASCADE,
    tokenHash VARCHAR(255),

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS refresh_tokens;
-- +goose StatementEnd
