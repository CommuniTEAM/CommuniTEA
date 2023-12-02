-- migrate:up
CREATE TABLE "users" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
);


-- migrate:down
DROP TABLE "users";

/*
This file should be deleted and a new migration file should be
generated when writing the actual users tables for the database,
because this test file was generated first and will thus migrate
first. Since our users tables depend on other tables (locations),
they need to migrate after the location tables to avoid errors.

Generate a new file in /CommuniTEA/backend with:
docker run --rm -it --network=host -v "$(pwd)/db:/db" ghcr.io/amacneil/dbmate new <NAME OF YOUR MIGRATION (ex. create_users_tables)>
*/
