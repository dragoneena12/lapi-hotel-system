
-- +migrate Up
CREATE TABLE IF NOT EXISTS stays (
id STRING PRIMARY KEY NOT NULL,
user_id STRING,
hotel_id STRING,
checkin DATETIME);

CREATE TABLE IF NOT EXISTS hotels (
id STRING PRIMARY KEY NOT NULL,
owner_id STRING,
name STRING,
location STRING,
key STRING,
carbon_awards STRING,
fullerene_awards STRING,
carbon_nanotube_awards STRING,
graphene_awards STRING,
diamond_awards STRING);

-- +migrate Down
DROP TABLE IF EXISTS stays;
DROP TABLE IF EXISTS hotels;
