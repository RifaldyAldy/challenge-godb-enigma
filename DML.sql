-- GET All mst_customer
 SELECT id,customer_name,phone_number FROM mst_customer;

-- INSERT mst_customer
 INSERT INTO mst_customer (customer_name,phone_number) VALUES ($1,$2);

-- UPDATE mst_customer 
 UPDATE mst_customer SET " + column + " = $1 WHERE id = $2;

-- DELETE mst_customer
 DELETE FROM mst_customer WHERE id = $1;
-- GET All mst_service
SELECT id, service_name, unit_name, price_per_unit FROM mst_service;

-- Insert mst_service
INSERT INTO mst_service (service_name,unit_name,price_per_unit) VALUES ($1,$2,$3);

-- Update mst_service
UPDATE mst_service SET " + column + " = $1 WHERE id = $2;

-- Delete mst_service
DELETE FROM mst_service WHERE id = $1;

-- GET transaksi_details join transaksi & mst_customer
SELECT c.customer_name,
	   td.received_by,
       s.service_name,
       t.quantity,
       s.unit_name,
       s.price_per_unit,
       SUM(s.price_per_unit * t.quantity) AS total_price
FROM transaksi AS t
JOIN mst_service AS s ON t.service_id = s.id
JOIN transaksi_details AS td ON td.id = t.transaction_id
JOIN mst_customer AS c ON c.id = td.customer_id
WHERE t.transaction_id = 3
GROUP BY c.customer_name,td.received_by, s.service_name, t.quantity, s.unit_name, s.price_per_unit;

-- Insert Transaksi tx
INSERT INTO transaksi_details (customer_id,entry_date,finish_date,received_by,total_price,quantity) VALUES ($1,$2,$3,$4,$5,$6) RETURNING ID;
INSERT into transaksi (transaction_id,service_id,quantity,total_price) VALUES ($1,$2,$3,$4);

-- Get data Transaksi 
SELECT td.id,c.customer_name,td.entry_date,td.finish_date,td.received_by FROM transaksi_details AS td JOIN mst_customer AS c ON td.customer_id = c.id

-- DELETE transaksi tx
DELETE FROM transaksi_details WHERE id = $1
DELETE FROM transaksi WHERE transaction_id = $1
