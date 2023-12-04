-- migrate:up
create table "users" (
    "id" serial primary key,
    "name" varchar(50) not null
);


-- migrate:down
drop table "users";

/*
This file should be deleted and a new migration file should be
generated when writing the actual users tables for the database,
because this test file was generated first and will thus migrate
first. Since our users tables depend on other tables (locations),
they need to migrate after the location tables to avoid errors.

Generate a new file in /CommuniTEA/backend with:
docker run --rm -it --network=host -v "$(pwd)/db:/db" ghcr.io/amacneil/dbmate new <NAME OF YOUR MIGRATION (ex. create_users_tables)>
*/
