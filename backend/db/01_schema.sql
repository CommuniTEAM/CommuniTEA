SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: business_followers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.business_followers (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    business_id uuid NOT NULL
);


--
-- Name: business_offered_teas; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.business_offered_teas (
    id uuid NOT NULL,
    business_id uuid NOT NULL,
    tea_id uuid NOT NULL
);


--
-- Name: business_reviews; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.business_reviews (
    id uuid NOT NULL,
    business uuid NOT NULL,
    author uuid NOT NULL,
    rating smallint NOT NULL,
    comment character varying(500),
    date date NOT NULL
);


--
-- Name: businesses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.businesses (
    id uuid NOT NULL,
    name character varying(100) NOT NULL,
    street_address character varying(100) NOT NULL,
    city uuid NOT NULL,
    zipcode character varying(5) NOT NULL,
    business_owner_id uuid NOT NULL
);


--
-- Name: event_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_categories (
    name character varying(50) NOT NULL
);


--
-- Name: event_category_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_category_tags (
    id uuid NOT NULL,
    event_id uuid NOT NULL,
    category character varying(50) NOT NULL
);


--
-- Name: event_cohost_permissions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_cohost_permissions (
    name character varying(50) NOT NULL
);


--
-- Name: event_cohosts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_cohosts (
    id uuid NOT NULL,
    event_id uuid NOT NULL,
    user_id uuid NOT NULL,
    permissions character varying(50) NOT NULL
);


--
-- Name: event_rsvps; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_rsvps (
    id uuid NOT NULL,
    event uuid NOT NULL,
    "user" uuid NOT NULL
);


--
-- Name: event_watchers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.event_watchers (
    id uuid NOT NULL,
    event_id uuid NOT NULL,
    user_id uuid NOT NULL
);


--
-- Name: events; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.events (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    host uuid NOT NULL,
    location_name character varying(100),
    street_address character varying(100) NOT NULL,
    city uuid NOT NULL,
    zipcode character varying(5) NOT NULL,
    start_time timestamp without time zone NOT NULL,
    end_time timestamp without time zone NOT NULL,
    html_description text,
    rsvps boolean NOT NULL,
    capacity integer,
    timezone_location character varying(50) NOT NULL,
    visible boolean DEFAULT false NOT NULL
);


--
-- Name: locations_cities; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.locations_cities (
    id uuid NOT NULL,
    name public.citext NOT NULL,
    state character varying(2) NOT NULL
);


--
-- Name: locations_states; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.locations_states (
    name character varying(50) NOT NULL,
    abbreviation character varying(2) NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: tea_aromatic_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tea_aromatic_tags (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    tea_id uuid NOT NULL
);


--
-- Name: tea_aromatics; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tea_aromatics (
    name character varying(50) NOT NULL
);


--
-- Name: tea_flavor_profile_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tea_flavor_profile_tags (
    id uuid NOT NULL,
    name character varying(50) NOT NULL,
    tea_id uuid NOT NULL
);


--
-- Name: tea_flavor_profiles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tea_flavor_profiles (
    name character varying(50) NOT NULL
);


--
-- Name: tea_origin_tags; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tea_origin_tags (
    id uuid NOT NULL,
    name character varying(50),
    tea_id uuid NOT NULL
);


--
-- Name: tea_origins; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tea_origins (
    name character varying(50) NOT NULL
);


--
-- Name: teas; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.teas (
    id uuid NOT NULL,
    name character varying(100) NOT NULL,
    img_url character varying(500),
    description text NOT NULL,
    brew_time character varying(50),
    brew_temp double precision,
    published boolean NOT NULL
);


--
-- Name: timezone_locations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.timezone_locations (
    zone_id character varying(50) NOT NULL
);


--
-- Name: user_favorite_teas; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_favorite_teas (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    tea_id uuid NOT NULL
);


--
-- Name: user_roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_roles (
    name character varying(25) NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    role character varying(25) NOT NULL,
    username public.citext NOT NULL,
    first_name character varying(50),
    last_name character varying(50),
    email character varying(200),
    password bytea NOT NULL,
    location uuid NOT NULL
);


--
-- Name: business_followers business_followers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_followers
    ADD CONSTRAINT business_followers_pkey PRIMARY KEY (id);


--
-- Name: business_offered_teas business_offered_teas_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_offered_teas
    ADD CONSTRAINT business_offered_teas_pkey PRIMARY KEY (id);


--
-- Name: business_reviews business_reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_reviews
    ADD CONSTRAINT business_reviews_pkey PRIMARY KEY (id);


--
-- Name: businesses businesses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_pkey PRIMARY KEY (id);


--
-- Name: event_categories event_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_categories
    ADD CONSTRAINT event_categories_pkey PRIMARY KEY (name);


--
-- Name: event_category_tags event_category_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_category_tags
    ADD CONSTRAINT event_category_tags_pkey PRIMARY KEY (id);


--
-- Name: event_cohost_permissions event_cohost_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_cohost_permissions
    ADD CONSTRAINT event_cohost_permissions_pkey PRIMARY KEY (name);


--
-- Name: event_cohosts event_cohosts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_cohosts
    ADD CONSTRAINT event_cohosts_pkey PRIMARY KEY (id);


--
-- Name: event_rsvps event_rsvps_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_rsvps
    ADD CONSTRAINT event_rsvps_pkey PRIMARY KEY (id);


--
-- Name: event_watchers event_watchers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_watchers
    ADD CONSTRAINT event_watchers_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: locations_cities locations_cities_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations_cities
    ADD CONSTRAINT locations_cities_pkey PRIMARY KEY (id);


--
-- Name: locations_states locations_states_name_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations_states
    ADD CONSTRAINT locations_states_name_key UNIQUE (name);


--
-- Name: locations_states locations_states_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations_states
    ADD CONSTRAINT locations_states_pkey PRIMARY KEY (abbreviation);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: tea_aromatic_tags tea_aromatic_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_aromatic_tags
    ADD CONSTRAINT tea_aromatic_tags_pkey PRIMARY KEY (id);


--
-- Name: tea_aromatics tea_aromatics_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_aromatics
    ADD CONSTRAINT tea_aromatics_pkey PRIMARY KEY (name);


--
-- Name: tea_flavor_profile_tags tea_flavor_profile_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_flavor_profile_tags
    ADD CONSTRAINT tea_flavor_profile_tags_pkey PRIMARY KEY (id);


--
-- Name: tea_flavor_profiles tea_flavor_profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_flavor_profiles
    ADD CONSTRAINT tea_flavor_profiles_pkey PRIMARY KEY (name);


--
-- Name: tea_origin_tags tea_origin_tags_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_origin_tags
    ADD CONSTRAINT tea_origin_tags_pkey PRIMARY KEY (id);


--
-- Name: tea_origins tea_origins_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_origins
    ADD CONSTRAINT tea_origins_pkey PRIMARY KEY (name);


--
-- Name: teas teas_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.teas
    ADD CONSTRAINT teas_pkey PRIMARY KEY (id);


--
-- Name: timezone_locations timezone_locations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.timezone_locations
    ADD CONSTRAINT timezone_locations_pkey PRIMARY KEY (zone_id);


--
-- Name: user_favorite_teas user_favorite_teas_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_favorite_teas
    ADD CONSTRAINT user_favorite_teas_pkey PRIMARY KEY (id);


--
-- Name: user_roles user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_roles
    ADD CONSTRAINT user_roles_pkey PRIMARY KEY (name);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: business_followers business_followers_business_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_followers
    ADD CONSTRAINT business_followers_business_id_fkey FOREIGN KEY (business_id) REFERENCES public.businesses(id);


--
-- Name: business_followers business_followers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_followers
    ADD CONSTRAINT business_followers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: business_offered_teas business_offered_teas_business_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_offered_teas
    ADD CONSTRAINT business_offered_teas_business_id_fkey FOREIGN KEY (business_id) REFERENCES public.businesses(id);


--
-- Name: business_offered_teas business_offered_teas_tea_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_offered_teas
    ADD CONSTRAINT business_offered_teas_tea_id_fkey FOREIGN KEY (tea_id) REFERENCES public.teas(id);


--
-- Name: business_reviews business_reviews_author_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_reviews
    ADD CONSTRAINT business_reviews_author_fkey FOREIGN KEY (author) REFERENCES public.users(id);


--
-- Name: business_reviews business_reviews_business_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.business_reviews
    ADD CONSTRAINT business_reviews_business_fkey FOREIGN KEY (business) REFERENCES public.businesses(id);


--
-- Name: businesses businesses_business_owner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_business_owner_id_fkey FOREIGN KEY (business_owner_id) REFERENCES public.users(id);


--
-- Name: businesses businesses_city_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.businesses
    ADD CONSTRAINT businesses_city_fkey FOREIGN KEY (city) REFERENCES public.locations_cities(id);


--
-- Name: event_category_tags event_category_tags_category_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_category_tags
    ADD CONSTRAINT event_category_tags_category_fkey FOREIGN KEY (category) REFERENCES public.event_categories(name);


--
-- Name: event_category_tags event_category_tags_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_category_tags
    ADD CONSTRAINT event_category_tags_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- Name: event_cohosts event_cohosts_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_cohosts
    ADD CONSTRAINT event_cohosts_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- Name: event_cohosts event_cohosts_permissions_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_cohosts
    ADD CONSTRAINT event_cohosts_permissions_fkey FOREIGN KEY (permissions) REFERENCES public.event_cohost_permissions(name);


--
-- Name: event_cohosts event_cohosts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_cohosts
    ADD CONSTRAINT event_cohosts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: event_rsvps event_rsvps_event_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_rsvps
    ADD CONSTRAINT event_rsvps_event_fkey FOREIGN KEY (event) REFERENCES public.events(id);


--
-- Name: event_rsvps event_rsvps_user_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_rsvps
    ADD CONSTRAINT event_rsvps_user_fkey FOREIGN KEY ("user") REFERENCES public.users(id);


--
-- Name: event_watchers event_watchers_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_watchers
    ADD CONSTRAINT event_watchers_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- Name: event_watchers event_watchers_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.event_watchers
    ADD CONSTRAINT event_watchers_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: events events_city_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_city_fkey FOREIGN KEY (city) REFERENCES public.locations_cities(id);


--
-- Name: events events_host_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_host_fkey FOREIGN KEY (host) REFERENCES public.businesses(id);


--
-- Name: events events_timezone_location_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_timezone_location_fkey FOREIGN KEY (timezone_location) REFERENCES public.timezone_locations(zone_id);


--
-- Name: locations_cities locations_cities_state_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.locations_cities
    ADD CONSTRAINT locations_cities_state_fkey FOREIGN KEY (state) REFERENCES public.locations_states(abbreviation);


--
-- Name: tea_aromatic_tags tea_aromatic_tags_name_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_aromatic_tags
    ADD CONSTRAINT tea_aromatic_tags_name_fkey FOREIGN KEY (name) REFERENCES public.tea_aromatics(name);


--
-- Name: tea_aromatic_tags tea_aromatic_tags_tea_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_aromatic_tags
    ADD CONSTRAINT tea_aromatic_tags_tea_id_fkey FOREIGN KEY (tea_id) REFERENCES public.teas(id);


--
-- Name: tea_flavor_profile_tags tea_flavor_profile_tags_name_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_flavor_profile_tags
    ADD CONSTRAINT tea_flavor_profile_tags_name_fkey FOREIGN KEY (name) REFERENCES public.tea_flavor_profiles(name);


--
-- Name: tea_flavor_profile_tags tea_flavor_profile_tags_tea_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_flavor_profile_tags
    ADD CONSTRAINT tea_flavor_profile_tags_tea_id_fkey FOREIGN KEY (tea_id) REFERENCES public.teas(id);


--
-- Name: tea_origin_tags tea_origin_tags_name_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_origin_tags
    ADD CONSTRAINT tea_origin_tags_name_fkey FOREIGN KEY (name) REFERENCES public.tea_origins(name);


--
-- Name: tea_origin_tags tea_origin_tags_tea_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.tea_origin_tags
    ADD CONSTRAINT tea_origin_tags_tea_id_fkey FOREIGN KEY (tea_id) REFERENCES public.teas(id);


--
-- Name: user_favorite_teas user_favorite_teas_tea_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_favorite_teas
    ADD CONSTRAINT user_favorite_teas_tea_id_fkey FOREIGN KEY (tea_id) REFERENCES public.teas(id);


--
-- Name: user_favorite_teas user_favorite_teas_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_favorite_teas
    ADD CONSTRAINT user_favorite_teas_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users users_location_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_location_fkey FOREIGN KEY (location) REFERENCES public.locations_cities(id);


--
-- Name: users users_role_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_fkey FOREIGN KEY (role) REFERENCES public.user_roles(name);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20231201040843'),
    ('20231201051016'),
    ('20231201090112'),
    ('20231205231215'),
    ('20231212055940'),
    ('20231212061145'),
    ('20231214012512'),
    ('20231214035807'),
    ('20240106002631'),
    ('20240318042602'),
    ('20240320042524');
