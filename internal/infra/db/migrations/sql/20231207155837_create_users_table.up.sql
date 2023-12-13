SET statement_timeout = 0;

--bun:split

create table public.users (
    id serial not null constraint users_pk primary key,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    name text not null,
    email text not null,
    phone text not null,
    password text not null,
    image text 
    );
    create unique index users_id_uindex on public.users (id)
    where deleted_at is null;

create unique index users_email_uindex on public.users (email)
where deleted_at is null;

create unique index users_phone_uindex on public.users (phone)
where deleted_at is null;
