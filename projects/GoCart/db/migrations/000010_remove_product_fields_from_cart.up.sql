DROP INDEX IF EXISTS idx_cart_product_id;

ALTER TABLE cart
    DROP COLUMN product_id,
    DROP COLUMN quantity;
