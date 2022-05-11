-- +goose Up
-- +goose StatementBegin
CREATE TABLE `transaction` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `sku` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `quantity` bigint NOT NULL,
  `order_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `subtotal` decimal(50,3) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `order_transaction_id_IDX` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE 'transaction';
-- +goose StatementEnd
