CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(255) not null,
    email varchar(255) not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE products
(
    id SERIAL PRIMARY KEY,
    user_id int references users(id) on delete cascade not null,
    name varchar(255) not null,
    price BIGINT not null,
    description varchar(255) not null,
    amount BIGINT not null,
    categort varchar(255) not null
);
CREATE TABLE orders
(
    id SERIAL PRIMARY KEY,
    user_id int references users(id) on delete cascade not null,
    product_id int references products(id) on delete cascade not null,
    product_price BIGINT references products(price) on delete cascade not null,
    product_name varchar(255) references products(name) on delete cascade not null
);
CREATE TABLE reviews
(
    id SERIAL PRIMARY KEY,
    user_id int references users(id) on delete cascade not null,
    product_id int references products(id) on delete cascade not null,
    text varchar(255) not null,
    evaluation BIGINT not null
);
