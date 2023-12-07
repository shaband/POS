SET statement_timeout = 0;

--bun:split

create table public.inovoices (
    id serial not null constraint inovoice_id primary key,
    created_at timestamp not null default now(),   
    updated_at timestamp not null default now(),   
    deleted_at timestamp,
    user_id integer REFERENCES users (id),
    inventory_id integer REFERENCES inventories (id),
    client_id integer REFERENCES clients (id),
    is_sell boolean not null default true,
    total_cost text not null ,
    total_price text not null 
);

--bun:split

create table public.inovoice_items (
    id serial not null constraint inovoice_items_id primary key,
    created_at timestamp not null default now(),   
    updated_at timestamp not null default now(),   
    deleted_at timestamp,
    inovoice_id integer REFERENCES inovoices (id),
    product_id integer REFERENCES products (id),
    user_id integer REFERENCES users (id),
    unit_sell_price text not null,
    unit_cost_price text not null,
    amount text not null ,
    total_cost text not null ,
    total_price text not null 
);

--bun:split
