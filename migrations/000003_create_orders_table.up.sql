-- ENUM untuk provider
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
        CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipped', 'delivered');
    END IF;
END
$$;

CREATE TABLE order_shippings (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    code VARCHAR(100),
    service VARCHAR(100),
    description TEXT,
    cost DOUBLE PRECISION,
    etd VARCHAR(50),
    address TEXT
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    order_number VARCHAR(255) UNIQUE NOT NULL,
    status order_status DEFAULT 'pending',
    total BIGINT NOT NULL,
    payment_method VARCHAR(50),
    billing_address TEXT,
    created_at TIMESTAMP NOT NULL,
    paid_at TIMESTAMP,
    shipped_at TIMESTAMP,
    delivered_at TIMESTAMP,
    shipping_id INTEGER NOT NULL REFERENCES order_shippings(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id)
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP,
    product_id UUID NOT NULL REFERENCES products(id),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE
);