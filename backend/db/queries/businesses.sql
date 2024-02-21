-- name: GetAllBusinesses :many
select * from businesses;

-- name: CreateBusiness :one
insert into businesses (
    id,
    name,
    street_address,
    city, state,
    zipcode,
    business_owner_id
)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;
