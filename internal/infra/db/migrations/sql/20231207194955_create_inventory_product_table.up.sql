SET statement_timeout = 0;

--bun:split

create  table public.inventory_product(
id serial not null constraint inventory_product_id primary key ,
product_id integer not null REFERENCES products,
inventory_id integer not null REFERENCES inventories,
amount integer not null default 0,
sell_invoices_count integer default 0,
buy_invoices_count integer default 0
);  