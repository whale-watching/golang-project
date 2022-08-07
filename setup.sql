DROP TABLE IF EXISTS customer;

-- Create a new table called 'customer'
CREATE TABLE customers (
  id serial primary key,
  firstname varchar(50),
  lastname varchar(50),
  email varchar(100) unique,
  password varchar(100),
  cc_customerid varchar(50),
  loggedin boolean,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  deleted_at timestamp 
);

-- Create a new table called 'orders'
CREATE TABLE orders (
  id serial primary key,
  customer_id int,
  product_id int,
  price int,
  purchase_date timestamp default current_timestamp,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  deleted_at timestamp
);

-- Create a new table called 'products'
CREATE TABLE products (
  id serial primary key,
  image varchar(100),
  imgalt varchar(100),
  description text,
  productname varchar(50),
  price float,
  promotion float,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  deleted_at timestamp
);