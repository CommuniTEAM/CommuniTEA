-- migrate:up
create table if not exists "teas" (
    "id" uuid primary key,
    "name" varchar(100) not null,
    "img_url" varchar(500),
    "description" text not null,
    "brew_time" varchar(50),
    "brew_temp" float,
    "published" boolean not null
);

create table if not exists "tea_flavor_profiles" (
    "name" varchar(50) primary key
);

create table if not exists "tea_flavor_profile_tags" (
    "id" uuid primary key,
    "name" varchar(50) not null references tea_flavor_profiles (name),
    "tea_id" uuid not null references teas (id)
);

create table if not exists "tea_origins" (
    "name" varchar(50) primary key
);

create table if not exists "tea_origin_tags" (
    "id" uuid primary key,
    "name" varchar(50) references tea_origins (name),
    "tea_id" uuid not null references teas (id)
);

create table if not exists "tea_aromatics" (
    "name" varchar(50) primary key
);

create table if not exists "tea_aromatic_tags" (
    "id" uuid primary key,
    "name" varchar(50) not null references tea_aromatics (name),
    "tea_id" uuid not null references teas (id)
);

-- migrate:down
drop table "teas";
drop table "tea_flavor_profiles";
drop table "tea_flavor_profile_tags";
drop table "tea_origins";
drop table "tea_origin_tags";
drop table "tea_aromatics";
drop table "tea_aromatic_tags";
