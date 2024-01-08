-- migrate:up
create table if not exists "locations_states" (
    "name" varchar(50) unique not null,
    "abbreviation" varchar(2) primary key
);

create table if not exists "locations_cities" (
    "id" uuid primary key,
    "name" citext not null,
    "state" varchar(2) not null references locations_states (abbreviation)
);

-- migrate:down
drop table "locations_states";
drop table "locations_cities";
