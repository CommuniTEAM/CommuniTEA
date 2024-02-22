--! Add any data needed for integration tests to this file

-- test suite user data
insert into "users" (
    "id",
    "location",
    "password",
    "role",
    "username"
) values (
    '372bcfb3-6b1d-4925-9f3d-c5ec683a4294',
    '4c33e0bc-3d43-4e77-aed0-b7aff09bb689',
    'hashed password',
    'user',
    'user'
);

insert into "users" (
    "id",
    "location",
    "password",
    "role",
    "username"
) values (
    '140e4411-a7f7-4c50-a2d4-f3d3fc9fc550',
    '4c33e0bc-3d43-4e77-aed0-b7aff09bb689',
    'hashed password',
    'business',
    'business'
);

insert into "users" (
    "id",
    "location",
    "password",
    "role",
    "username"
) values (
    'e6473137-f4ef-46cc-a5e5-96ccb9d41043',
    '4c33e0bc-3d43-4e77-aed0-b7aff09bb689',
    'hashed password',
    'admin',
    'admin'
);

-- locations testdata
insert into "locations_cities" (
    "id",
    "name",
    "state"
) values (
    'b6df94c1-4d68-4a1e-b702-60f5cabcebcc',
    'Tacoma',
    'WA'
);

insert into "locations_cities" (
    "id",
    "name",
    "state"
) values (
    '6937755c-7e87-4226-9692-36d3019be32a',
    'New York',
    'NY'
);

insert into "locations_cities" (
    "id",
    "name",
    "state"
) values (
    '07eca16a-8ee1-4c1a-831e-cb984a851bf3',
    'Kansas City',
    'KS'
);

-- teas testdata
insert into "teas" (
    "id",
    "name",
    "img_url",
    "description",
    "brew_time",
    "brew_temp",
    "published"
) values (
    'c64ff5ab-7323-4142-9077-aea320c3c4cc',
    'Earl Grey',
    'https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
    'It is a black tea mix',
    '20 minutes',
    175,
    false
);

insert into "teas" (
    "id",
    "name",
    "img_url",
    "description",
    "brew_time",
    "brew_temp",
    "published"
) values (
    'c64ff5ab-7323-4142-9077-aea320c3c4cf',
    'Black Tea',
    'https://images.unsplash.com/photo-1605618826115-fb9e775cfb40?q=80&w=1779&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
    'It is a black tea mix',
    '20 minutes',
    175,
    false
);

-- business test data
insert into "businesses" (
    "id",
    "name",
    "street_address",
    "city",
    "state",
    "zipcode",
    "business_owner_id"
) values (
    'e6e8e3e3-3e3e-4e3e-8e3e-3e3e3e3e3e3e',
    'Test Business',
    '123 Test St',
    '07eca16a-8ee1-4c1a-831e-cb984a851bf3',
    'KS',
    '98444',
    '140e4411-a7f7-4c50-a2d4-f3d3fc9fc550'
);

-- events testdata
insert into "events" (
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
) values (
    'e6e8e3e3-3e3e-4e3e-8e3e-3e3e3e3e3e3e',
    'Test Event',
    'e6e8e3e3-3e3e-4e3e-8e3e-3e3e3e3e3e3e',
    'Test Location',
    '123 Test St',
    '07eca16a-8ee1-4c1a-831e-cb984a851bf3',
    'WA',
    '98444',
    '2021-12-31',
    '12:00:00',
    '14:00:00',
    'Test Description',
    'Test Description',
    true,
    100
);
