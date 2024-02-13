-- name: GetTeas :many
select *
from teas
where published = $1;

-- name: GetTea :one
select *
from teas
where id = $1 limit 1;

-- name: CreateTea :one
insert into teas (
    id,
    name,
    img_url,
    description,
    brew_time,
    brew_temp,
    published
)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

-- name: UpdateTea :one
update teas
set
    "name" = $2,
    "img_url" = $3,
    "description" = $4,
    "brew_time" = $5,
    "brew_temp" = $6,
    "published" = $7
where "id" = $1
returning *;

-- name: DeleteTea :exec
delete from teas
where id = $1;

-- name: CreateTeaOrigins :one
insert into tea_origins (name)
values ($1)
returning *;

-- name: GetAllTeaOrigins :many
select * from tea_origins;

-- name: GetTeaOrigin :one
select *
from tea_origins
where name = $1 limit 1;

-- name: DeleteTeaOrigins :exec
delete from tea_origins
where name = $1;

-- name: GetAllTeaAromatics :many
select * from tea_aromatics;

-- name: GetTeaAromatic :one
select *
from tea_aromatics
where name = $1 limit 1;

-- name: CreateTeaAromatics :one
insert into tea_aromatics (name)
values ($1)
returning *;

-- name: DeleteTeaAromatics :exec
delete from tea_aromatics
where name = $1;

-- name: GetAllTeaFlavorProfiles :many
select * from tea_flavor_profiles;

-- name: GetTeaFlavorProfile :one
select *
from tea_flavor_profiles
where name = $1 limit 1;

-- name: CreateTeaFlavorProfiles :one
insert into tea_flavor_profiles (name)
values ($1)
returning *;

-- name: DeleteTeaFlavorProfiles :exec
delete from tea_flavor_profiles
where name = $1;

-- name: CreateTeaOriginTags :one
insert into tea_origin_tags (id, name, tea_id)
values ($1, $2, $3)
returning *;

-- name: GetTeaOriginTags :many
select tags.name
from tea_origin_tags as tags
inner join teas on tags.tea_id = teas.id
where tags.tea_id = $1;

-- name: GetTeaOriginTag :one
select *
from tea_origin_tags
where id = $1;

-- name: UpdateTeaOriginTag :exec
update tea_origin_tags
set name = $2
where id = $1;

-- name: DeleteTeaOriginTag :exec
delete from tea_origin_tags
where id = $1;

-- name: CreateTeaAromaticTags :one
insert into tea_aromatic_tags (id, name, tea_id)
values ($1, $2, $3)
returning *;

-- name: GetTeaAromaticTags :many
select tags.name
from tea_aromatic_tags as tags
inner join teas on tags.tea_id = teas.id
where tags.tea_id = $1;

-- name: GetTeaAromaticTag :one
select *
from tea_aromatic_tags
where id = $1;

-- name: UpdateTeaAromaticTag :exec
update tea_aromatic_tags
set name = $2
where id = $1;

-- name: DeleteTeaAromaticTag :exec
delete from tea_aromatic_tags
where id = $1;

-- name: CreateTeaFlavorProfileTags :one
insert into tea_flavor_profile_tags (id, name, tea_id)
values ($1, $2, $3)
returning *;

-- name: GetTeaFlavorProfileTags :many
select tags.name
from tea_flavor_profile_tags as tags
inner join teas on tags.tea_id = teas.id
where tags.tea_id = $1;

-- name: GetTeaFlavorProfileTag :one
select *
from tea_flavor_profile_tags
where id = $1;

-- name: UpdateTeaFlavorProfileTag :exec
update tea_flavor_profile_tags
set name = $2
where id = $1;

-- name: DeleteTeaFlavorProfileTag :exec
delete from tea_flavor_profile_tags
where id = $1;
