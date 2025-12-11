ALTER TABLE carts
ADD CONSTRAINT uq_user_product UNIQUE (user_id, product_id);
