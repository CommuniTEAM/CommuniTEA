-- migrate:up
insert into "user_roles"
values ('user')
on conflict (name) do nothing;

insert into "user_roles"
values ('business')
on conflict (name) do nothing;

insert into "user_roles"
values ('admin')
on conflict (name) do nothing;

-- migrate:down
