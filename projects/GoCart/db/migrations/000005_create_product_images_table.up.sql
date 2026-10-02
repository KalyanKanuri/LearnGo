CREATE TABLE
    product_images (
        id SERIAL PRIMARY KEY,
        product_id INT NOT NULL,
        image_url TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL
    );

CREATE INDEX idx_product_images_product_id ON product_images (product_id);

CREATE INDEX idx_product_images_image_url ON product_images (image_url);