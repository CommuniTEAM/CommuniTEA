-- name: CreateCity :one
insert into locations_cities
values ($1, $2, $3)
returning *;

-- name: GetAllCities :many
select * from locations_cities;

-- ! THIS IS A DEBUG QUERY: DELETE FOR PROD
-- name: GetCity :one
select "id" from locations_cities
where ("name" = 'string')
limit 1;

-- name: test :many
select * from locations_states;
