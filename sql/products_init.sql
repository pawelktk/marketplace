CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  description VARCHAR NOT NULL,
  stock INTEGER NOT NULL,
  price FLOAT NOT NULL
);

-- INSERT INTO products(name, description, stock, price) VALUES('Kawa', 'Dobra kawa', 100, 12);
