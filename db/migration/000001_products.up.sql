CREATE TABLE products (
  id varchar(36) PRIMARY KEY,
  title varchar NOT NULL,
  description varchar NOT NULL,
  price int NOT NULL,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON products(id);
