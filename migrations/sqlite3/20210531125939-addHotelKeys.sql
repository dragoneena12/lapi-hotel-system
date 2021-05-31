
-- +migrate Up
ALTER TABLE hotels ADD COLUMN hotelKey STRING;

-- +migrate Down
