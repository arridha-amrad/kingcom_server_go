ALTER TABLE order_items
    DROP COLUMN IF EXISTS price_at_order,
    DROP COLUMN IF EXISTS discount_at_order,
    DROP COLUMN IF EXISTS final_price_at_order;
