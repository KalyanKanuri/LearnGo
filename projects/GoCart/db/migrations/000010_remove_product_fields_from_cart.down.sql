ALTER TABLE cart
    ADD COLUMN product_id INT NOT NULL DEFAULT 0,
    ADD COLUMN quantity INT NOT NULL DEFAULT 0;

ALTER TABLE cart
    ALTER COLUMN product_id DROP DEFAULT,
    ALTER COLUMN quantity DROP DEFAULT;

CREATE INDEX idx_cart_product_id ON cart (product_id);
