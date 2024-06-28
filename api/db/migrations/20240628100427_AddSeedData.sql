-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied.

-- Insert companies
INSERT INTO companies (id, name) VALUES
  ('d290f1ee-6c54-4b01-90e6-d701748f0851', 'Google'),
  ('4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'New Relic');

-- Insert customers associated with Google
INSERT INTO customers (id, company_id, first_name, last_name) VALUES
  ('11a611b8-5bf0-4f61-9e21-94ec7e084f3a', 'd290f1ee-6c54-4b01-90e6-d701748f0851', 'John', 'Doe'),
  ('24b7bc2e-606c-44f0-b5f3-2e02e3548d52', 'd290f1ee-6c54-4b01-90e6-d701748f0851', 'Jane', 'Doe'),
  ('3c6e80b7-7576-4c8d-82c3-6d5b7e5192fc', 'd290f1ee-6c54-4b01-90e6-d701748f0851', 'Mike', 'Smith');

-- Insert customers associated with New Relic
INSERT INTO customers (id, company_id, first_name, last_name) VALUES
  ('49d9e0b3-481f-47e0-8b7d-5d2e1d4760f1', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'Alice', 'Johnson'),
  ('50f5a2d3-6c52-4d2e-8e39-7d6c1e676b2a', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'Bob', 'Brown'),
  ('62e9f7b2-7609-4f3d-9c6b-6e1f5b5f8e6c', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'Carol', 'Williams'),
  ('74f4e2b8-819c-4e1b-8d7d-7c5f5b8f6e6f', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'David', 'Miller'),
  ('85d5c3e9-91a2-4f3f-9e2d-8f7e6c7d8d6e', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'Eve', 'Davis'),
  ('96e6e0b2-a2c0-4d2e-9e2d-9f6e1f7e8f7e', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'Frank', 'Garcia'),
  ('a7f7e1b3-b2d1-4e1d-9f7e-a0c1f7e9e0f8', '4c9fc733-b7a4-46ee-9e72-1af8889dbff2', 'Grace', 'Martinez');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back.

-- Delete customers
DELETE FROM customers WHERE company_id IN (
  'd290f1ee-6c54-4b01-90e6-d701748f0851',
  '4c9fc733-b7a4-46ee-9e72-1af8889dbff2'
);

-- Delete companies
DELETE FROM companies WHERE id IN (
  'd290f1ee-6c54-4b01-90e6-d701748f0851',
  '4c9fc733-b7a4-46ee-9e72-1af8889dbff2'
);
