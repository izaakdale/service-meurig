CREATE TABLE order_items (
  id varchar PRIMARY KEY,
  order_id varchar(36) NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  product_id varchar(36) NOT NULL,
  FOREIGN KEY (product_id) REFERENCES products(id),
  product_name varchar NOT NULL,
  quantity int NOT NULL
);

CREATE INDEX ON order_items(id);

-- Add an index to the order_id column for faster lookups
CREATE INDEX order_items_order_id_idx ON order_items ("order_id");