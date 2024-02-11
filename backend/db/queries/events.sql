-- name: GetEventByID :one
select
    "id",
    "name",
    "host",
    "location_name",
    "street_address",
    "city",
    "state",
    "zipcode",
    "date",
    "start_time",
    "end_time",
    "md_description",
    "html_description",
    "rsvps",
    "capacity"
from events
where "id" = $1;