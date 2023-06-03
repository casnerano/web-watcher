create extension if not exists "uuid-ossp";

create table if not exists users (
    uuid uuid primary key default uuid_generate_v4() not null,
    email character varying(100) not null,
    password character varying(100) not null,
    name character varying(100),
    last_name character varying(100),
    created_at timestamp default now() not null,
    constraint users_unique_email unique (email)
);
