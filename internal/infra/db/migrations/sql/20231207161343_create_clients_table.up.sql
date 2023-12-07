SET statement_timeout = 0;

--bun:split

create table public.clients(
    id  serial not null constraint clients_pk primary key,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp,
    name text not null,
    email text not null,
    phone text not null,
    image text ,
    is_supplier  boolean default false,
    is_client  boolean default true
);

create unique index clients_id_uindex on public.clients (id)
where deleted_at is null;

create unique index clients_email_uindex on public.clients (email)
where deleted_at is null;

create unique index clients_phone_uindex on public.clients (phone)
where deleted_at is null;

create index clients_is_supplier_index on public.clients (is_supplier);
create index clients_is_client_index on public.clients (is_client);
