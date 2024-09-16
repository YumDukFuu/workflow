-- +goose Up
ALTER TABLE items ADD owner_id integer ;

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE items
DROP COLUMN owner_id;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
