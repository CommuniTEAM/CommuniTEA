-- name: CreateCity :one
insert into locations_cities
values ($1, $2, $3)
returning *;

-- name: GetAllCities :many
select * from locations_cities;
