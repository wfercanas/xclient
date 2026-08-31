CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  tenant_id INT REFERENCES tenants(id),
  customer_type TEXT NOT NULL CHECK (customer_type IN ('person', 'company')),
  customer_id TEXT NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL CHECK (status IN ('active', 'inactive', 'retiring', 'retired')),
  association_date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (tenant_id, customer_id)
);