-- migrate:up
insert into "locations_cities"
values ('4c33e0bc-3d43-4e77-aed0-b7aff09bb689', 'Seattle', 'WA')
on conflict (id) do nothing;

-- migrate:down
