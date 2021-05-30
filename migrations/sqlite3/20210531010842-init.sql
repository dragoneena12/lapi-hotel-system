
-- +migrate Up
CREATE TABLE IF NOT EXISTS stays (
id INTEGER PRIMARY KEY NOT NULL,
hotel_id STRING,
checkin DATETIME,
checkout DATETIME,
user STRING);

CREATE TABLE IF NOT EXISTS hotels (
id INTEGER PRIMARY KEY NOT NULL,
name STRING,
location STRING,
owner STRING,
carbonAwards STRING,
fullereneAwards STRING,
carbonNanotubeAwards STRING,
grapheneAwards STRING,
diamondAwards STRING);

-- +migrate Down
DROP TABLE IF EXISTS stays;
DROP TABLE IF EXISTS hotels;