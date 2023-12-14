-- migrate:up
create table "events" (
    "id" uuid primary key,
    "name" varchar(50) not null,
    "host" uuid not null references businesses (id),
    "location_name" varchar(100),
    "street_address" varchar(100) not null,
    "city" uuid not null references locations_cities (id),
    "state" varchar(2) not null references locations_states (abbreviation),
    "zipcode" varchar(5) not null,
    "date" date not null,
    "start_time" time not null,
    "end_time" time not null,
    "description" text not null,
    "headliner_one" varchar(250),
    "headliner_two" varchar(250),
    "highlight_one" varchar(250),
    "highlight_two" varchar(250),
    "highlight_three" varchar(250),
    "rsvps" boolean not null,
    "capacity" int
);

create table "event_cohost_permissions" (
    "name" varchar(50) primary key
);

create table "event_cohosts" (
    "id" uuid primary key,
    "event_id" uuid not null references events (id),
    "user_id" uuid not null references users (id),
    "permissions" varchar(50) not null references event_cohost_permissions (name)
);

create table "event_watchers" (
    "id" uuid primary key,
    "event_id" uuid not null references events (id),
    "user_id" uuid not null references users (id)
);

create table "event_categories" (
    "name" varchar(50) primary key
);

create table "event_category_tags" (
    "id" uuid primary key,
    "event_id" uuid not null references events (id),
    "category" varchar(50) not null references event_categories (name)
);

create table "event_rsvps" (
    "id" uuid primary key,
    "event" uuid not null references events (id),
    "user" uuid not null references users (id)
);


-- migrate:down
drop table "events";
drop table "event_cohosts";
drop table "event_cohost_permissions";
drop table "event_watches";
drop table "event_category_tags";
drop table "event_categories";
drop table "event_rsvps";
