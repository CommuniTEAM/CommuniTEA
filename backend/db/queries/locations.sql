-- name: GetAllStates :many
select * from locations_states;

-- name: CreateCity :one
insert into locations_cities ("id", "name", "state")
values ($1, $2, $3)
returning *;

-- name: GetAllCities :many
select * from locations_cities;

-- name: GetAllCitiesInState :many
select * from locations_cities
where "state" = $1;

-- name: GetCity :one
select
    "name",
    "state"
from locations_cities
where "id" = $1;

-- name: GetCityID :one
select "id"
from locations_cities
where "name" = $1
    and "state" = $2;

-- name: UpdateCityName :one
update locations_cities
set "name" = $1
where "id" = $2
returning *;

-- name: DeleteCity :exec
delete from locations_cities
where "id" = $1;
