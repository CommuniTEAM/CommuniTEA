-- migrate:up
create table if not exists "user_roles" (
    "name" varchar(25) primary key
);

create table if not exists "users" (
    "id" uuid primary key,
    "role" varchar(25) not null references user_roles (name),
    "username" varchar(50) unique not null,
    "first_name" varchar(50),
    "last_name" varchar(50),
    "email" varchar(200) unique,
    "password" varchar(1000) not null,
    "location" uuid not null references locations_cities (id)
);

create table if not exists "user_favorite_teas" (
    "id" uuid primary key,
    "user_id" uuid not null references users (id),
    "tea_id" uuid not null references teas (id)
);

-- migrate:down
drop table "user_favorite_teas";
drop table "users";
drop table "user_roles";
