-- migrate:up
create table if not exists "businesses" (
    "id" uuid primary key,
    "name" varchar(100) not null,
    "street_address" varchar(100) not null,
    "city" uuid not null references locations_cities (id),
    "state" varchar(2) not null references locations_states (abbreviation),
    "zipcode" varchar(5) not null,
    "business_owner_id" uuid not null references users (id)
);

create table if not exists "business_offered_teas" (
    "id" uuid primary key,
    "business_id" uuid not null references businesses (id),
    "tea_id" uuid not null references teas (id)
);

create table if not exists "business_followers" (
    "id" uuid primary key,
    "user_id" uuid not null references users (id),
    "business_id" uuid not null references businesses (id)
);

create table if not exists "business_reviews" (
    "id" uuid primary key,
    "business" uuid not null references businesses (id),
    "author" uuid not null references users (id),
    "rating" smallint not null,
    "comment" varchar(500),
    "date" date not null
);

-- migrate:down
drop table "businesses";
drop table "business_offered_teas";
drop table "business_followers";
drop table "business_reviews";
