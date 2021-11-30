-- migrate:up
create extension if not exists "uuid-ossp";

create table users (
    id uuid default uuid_generate_v4() primary key,
    username text not null unique,
    email text not null unique,
    password_hash text not null,
    full_name text,
    bio text,
    avatar_url text,
    -- email_verified
    -- active
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
-- migrate:down
drop table users;

drop extension "uuid-ossp";