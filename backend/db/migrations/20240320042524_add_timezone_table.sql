-- migrate:up
create table if not exists "timezone_locations" (
    "zone_id" varchar(50) primary key
);

insert into "timezone_locations"
values ('America/Adak')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Anchorage')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Boise')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Denver')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Detroit')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Indianapolis')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Knox')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Marengo')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Petersburg')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Tell_City')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Vevay')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Vincennes')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Indiana/Winamac')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Juneau')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Kentucky/Louisville')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Kentucky/Monticello')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Los_Angeles')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Menominee')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Metlakatla')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/New_York')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Nome')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/North_Dakota/Beulah')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/North_Dakota/Center')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/North_Dakota/New_Salem')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Phoenix')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Sitka')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('America/Yakutat')
on conflict (zone_id) do nothing;

insert into "timezone_locations"
values ('Pacific/Honolulu')
on conflict (zone_id) do nothing;

alter table "events"
add if not exists "timezone_location" varchar(50);

update "events"
set "timezone_location" = 'America/Los_Angeles';

alter table "events"
add constraint events_timezone_location_fkey foreign key ("timezone_location") references "timezone_locations" ("zone_id");

alter table "events"
alter "timezone_location" set not null;


-- migrate:down
alter table "events" drop "timezone_location";

drop table "timezone_locations";
