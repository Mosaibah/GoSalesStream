CREATE DATABASE IF NOT EXISTS "GoSalesStream";

SET DATABASE  = "GoSalesStream";

CREATE SEQUENCE transactions_seq
  start 1
  increment 1;

CREATE SEQUENCE customers_seq
  start 1
  increment 1;
 
CREATE SEQUENCE products_seq
  start 1
  increment 1;
 
CREATE TABLE customers (
	"id" int NOT null default nextval('customers_seq'),
	"name" VARCHAR(255) NOT NULL,
	CONSTRAINT "customers_pk" PRIMARY KEY ("id")
);

CREATE TABLE products (
	"id" int NOT null default nextval('products_seq'),
	"name" VARCHAR(255) NOT NULL,
	CONSTRAINT "products_pk" PRIMARY KEY ("id")
);

CREATE TABLE transactions (
	"id" int NOT null default nextval('transactions_seq'),
	"customer_id" integer NOT NULL,
	"product_id" integer NOT NULL,
	"price" integer NOT NULL,
	"quantity" integer not null,
	"created_at" TIMESTAMP NOT null DEFAULT now(),
	CONSTRAINT "transactions_pk" PRIMARY KEY ("id")
); 

ALTER TABLE transactions ADD CONSTRAINT "transactions_fk0" FOREIGN KEY ("customer_id") REFERENCES "customers"("id");
ALTER TABLE transactions ADD CONSTRAINT "transactions_fk1" FOREIGN KEY ("product_id") REFERENCES "products"("id");

insert into customers (name) values('Abdulrahman');
insert into customers (name) values('Ibramin');
insert into customers (name) values('Omar');
insert into customers (name) values('Mohammed');
insert into customers (name) values('Abdulaziz');
insert into customers (name) values('Majed');
insert into products (name) values('mac');
insert into products (name) values('book');

insert into transactions (customer_id, product_id, quantity, price) values (1,1,15, 99);
insert into transactions (customer_id, product_id, quantity, price) values (2,2,10, 230);
insert into transactions (customer_id, product_id, quantity, price) values (2,2,10, 120);
insert into transactions (customer_id, product_id, quantity, price) values (3,1,10, 9812);
insert into transactions (customer_id, product_id, quantity, price) values (3,2,10, 343);
insert into transactions (customer_id, product_id, quantity, price) values (3,2,10, 3);
insert into transactions (customer_id, product_id, quantity, price) values (3,1,10, 9872);
insert into transactions (customer_id, product_id, quantity, price) values (4,2,10, 23);
insert into transactions (customer_id, product_id, quantity, price) values (5,1,10, 35);
insert into transactions (customer_id, product_id, quantity, price) values (6,2,10, 645);