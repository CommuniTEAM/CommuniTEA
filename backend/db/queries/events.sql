-- name: GetEventByID :one
select
    "id",
    "name",
    "host",
    "location_name",
    "street_address",
    "city",
    "zipcode",
    "start_time",
    "end_time",
    "md_description",
    "html_description",
    "rsvps",
    "capacity",
    "timezone_location"
from events
where "id" = $1;


-- name: CreateEvent :one
insert into "events" ("id", "name", "host", "location_name", "street_address", "city", "zipcode", "start_time", "end_time", "md_description", "html_description", "rsvps", "capacity", "timezone_location")
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
returning *;
