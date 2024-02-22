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

-- name: UpdateBusiness :one
update businesses
set
    "name" = $2,
    "street_address" = $3,
    "city" = $4,
    "state" = $5,
    "zipcode" = $6,
    "business_owner_id" = $7
where "id" = $1
returning *;

-- name: DeleteBusiness :exec
delete from businesses
where id = $1;
