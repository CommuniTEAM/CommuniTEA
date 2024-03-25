-- migrate:up
alter table "events"
alter "start_time" type timestamp;

alter table "events"
alter "end_time" type timestamp;

alter table "events"
add if not exists "visible" boolean not null default false;

alter table "events"
drop if exists "date";

alter table "events"
drop if exists "md_description";

alter table "events"
drop if exists "state";

alter table "businesses"
drop if exists "state";

-- migrate:down
