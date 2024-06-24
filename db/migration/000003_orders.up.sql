CREATE TABLE orders (
  id varchar PRIMARY KEY,
  customer_id varchar(36) NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
  status varchar(255) NOT NULL DEFAULT 'cart',
  total int NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON orders(id);

-- Add an index to the customer_id column for faster lookups
CREATE INDEX orders_customer_id_idx ON orders ("customer_id");

