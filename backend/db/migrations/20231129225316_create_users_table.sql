-- migrate:up
CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
);


-- migrate:down
DROP TABLE "users";
