
-- +migrate Up
ALTER TABLE hotels ADD hotelKey varchar(255);

-- +migrate Down
ALTER TABLE hotels DROP hotelKey;