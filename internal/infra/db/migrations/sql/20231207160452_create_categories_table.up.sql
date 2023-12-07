SET statement_timeout = 0;

--bun:split

create table public.categories (
    id serial not null constraint categories_pk primary key,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    name text not null,
    categories_id integer  REFERENCES categories
);

create unique index categories_id_uindex on public.categories (id)
where deleted_at is null;

