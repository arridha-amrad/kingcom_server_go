CREATE TABLE products (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL,
    weight DOUBLE PRECISION NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    description TEXT NOT NULL,
    specification TEXT,
    stock INTEGER NOT NULL DEFAULT 0,
    video_url VARCHAR(500),
    discount INTEGER,
    CONSTRAINT chk_discount CHECK (discount >= 0)
);

CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    url VARCHAR(500) NOT NULL,
    product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE
);
