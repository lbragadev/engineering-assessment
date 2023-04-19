
-- +migrate Up
CREATE TABLE IF NOT EXISTS food_trucks
(
    id        serial primary key NOT NULL,
    location_id        integer NOT NULL,
    name            varchar(100) NOT NULL,
    address            varchar(100) NOT NULL,
    status            varchar(50),
    facility_type            varchar(50),
    location_description            varchar(200),
    food_items            varchar(500),
    latitude            numeric,
    longitude            numeric,
    created_at timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS food_trucks;
