CREATE TYPE order_status AS ENUM (
    'Pending',
    'Confirmed',
    'Shipped',
    'Delivered',
    'Completed',
    'Cancelled'
);

CREATE TABLE
    orders (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES users (id),
        status order_status NOT NULL,
        total_price NUMERIC(10, 2) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL
    );

CREATE INDEX idx_orders_user_id ON orders (user_id);

CREATE INDEX idx_orders_status ON orders (status);

CREATE INDEX idx_orders_total_price ON orders (total_price);