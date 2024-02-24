-- Table: mst_customer
CREATE TABLE mst_customer (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(15) NOT NULL
);
-- Table: mst_service
CREATE TABLE mst_service (
    id SERIAL PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    unit_name VARCHAR(10) NOT NULL,
    price_per_unit NUMERIC(100) NOT NULL
);

-- Table: transaksi_details
CREATE TABLE transaksi_details (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES mst_customer(id),
    entry_date DATE NOT NULL,
    finish_date DATE NOT NULL,
    received_by VARCHAR(255) NOT NULL,
    total_price NUMERIC(100) NOT NULL,
    quantity INT NOT NULL
);

-- Table: transaksi
CREATE TABLE transaksi(
    id SERIAL PRIMARY KEY,
    transaction_id INT,
    service_id INT,
    quantity INT,
    total_price DECIMAL(10,2),
    FOREIGN KEY (transaction_id) REFERENCES transaksi_details(id),
    FOREIGN KEY (service_id) REFERENCES mst_service(id)
);

