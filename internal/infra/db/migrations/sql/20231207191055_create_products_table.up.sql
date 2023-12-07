SET statement_timeout = 0;

--bun:split

create table products (
id serial not null constraint product_id primary key,
created_at timestamp not null default now(),
updated_at timestamp not null default now(),
deleted_at timestamp,   
name  text not null, 
code  text not null, 
cost_price text not null,
sell_price text not null,
image text,
category_id  integer REFERENCES categories (id)
);

create unique index products_code_uindex on public.products (code)
where deleted_at is null;


create unique index products_id_uindex on public.products (id)
where deleted_at is null;
