CREATE TABLE customers (
  id varchar(36) PRIMARY KEY,
  username VARCHAR(100) NOT NULL
);

CREATE INDEX ON customers(id);
