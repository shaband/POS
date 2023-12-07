SET statement_timeout = 0;

--bun:split


create table inventories(

id serial not null constraint inventory_id  primary key,
created_at timestamp not null default now(),
updated_at timestamp not null default now(),
deleted_at timestamp ,
name text not null,
address text 
);

create unique index inventory_uindex_name on public.inventories (name) 
where deleted_at is null;
