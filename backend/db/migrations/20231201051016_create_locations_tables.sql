-- migrate:up
CREATE TABLE IF NOT EXISTS "locations_states" (
    "name" VARCHAR(50) UNIQUE NOT NULL,
    "abbreviation" VARCHAR(2) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS "locations_cities" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL,
    "state" VARCHAR(2) NOT NULL REFERENCES locations_states (abbreviation)
);

-- migrate:down
DROP TABLE "locations_states";
DROP TABLE "locations_cities";
