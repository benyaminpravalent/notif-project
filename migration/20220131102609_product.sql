-- +goose Up
-- +goose StatementBegin
CREATE TABLE `product` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `sku` varchar(200) COLLATE utf8mb4_general_ci NOT NULL,
  `brand_id` bigint NOT NULL,
  `stock` bigint NOT NULL DEFAULT '0',
  `price` decimal(50,3) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `product_FK` (`brand_id`),
  CONSTRAINT `product_FK` FOREIGN KEY (`brand_id`) REFERENCES `brand` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO product (sku,brand_id,stock,price) VALUES ('sku_product-2',1,10,50000.0);
INSERT INTO product (sku,brand_id,stock,price) VALUES ('sku_product-1',1,3,10000.0);
INSERT INTO product (sku,brand_id,stock,price) VALUES ('sku_product-3',1,5,30000.0);
INSERT INTO product (sku,brand_id,stock,price) VALUES ('sku_product-4',1,21,100000.0);
INSERT INTO product (sku,brand_id,stock,price) VALUES ('sku_product-5',1,12,500000.0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE 'product';
-- +goose StatementEnd
