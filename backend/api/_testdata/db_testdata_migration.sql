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
