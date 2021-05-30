
-- +migrate Up
CREATE TABLE IF NOT EXISTS stays (
id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
hotel_id VARCHAR(255),
checkin DATETIME,
checkout DATETIME,
user VARCHAR(255));

CREATE TABLE IF NOT EXISTS hotels (
id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
name VARCHAR(255),
location VARCHAR(255),
owner VARCHAR(255),
carbonAwards VARCHAR(255),
fullereneAwards VARCHAR(255),
carbonNanotubeAwards VARCHAR(255),
grapheneAwards VARCHAR(255),
diamondAwards VARCHAR(255));

-- +migrate Down
DROP TABLE IF EXISTS stays;
DROP TABLE IF EXISTS hotels;