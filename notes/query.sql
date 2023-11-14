-- membuat database dengan nama pijarcamp
CREATE DATABASE pijarcamp;

-- membuat tabel products
CREATE TABLE products (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255) NOT NULL,
  price INTEGER NOT NULL,
  qty INTEGER NOT NULL
);

-- menambahkan data baru ke tabel products
INSERT INTO products (name, description, price, qty)
VALUES (?, ?, ?, ?);

-- mendapatkan semua data dari tabel products
SELECT id, name, description, price, qty
FROM products
ORDER BY name ASC;

-- memperbaharui salah satu data pada tabel products
UPDATE products
SET name = ?, description = ?, price = ?, qty = ?
WHERE id = ?;

-- menghapus salah satu data pada tabel products
DELETE FROM products
WHERE id = ?;
