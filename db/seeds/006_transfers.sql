CREATE TEMP TABLE mock_transfers (
  to_account_id BIGINT,
  from_account_id BIGINT,
  transaction_type VARCHAR(30),
  transaction_description TEXT,
  transaction_status VARCHAR(30),
  currency VARCHAR(3),
  amount DECIMAL(19,4)
);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (52, 50, 'transfer', 'Online transfer', 'failed', 'CHF', 4201.54);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (8, 145, 'transfer', 'Transfer to savings', 'cancelled', 'INR', 3779.27);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (5, 82, 'transfer', 'Internal transfer', 'completed', 'CHF', 1067.36);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (3, 75, 'transfer', 'Online transfer', 'completed', 'CHF', 3218.02);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (102, 82, 'transfer', 'Internal transfer', 'completed', 'EUR', 3897.95);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (119, 59, 'transfer', 'Transfer to savings', 'pending', 'EUR', 327.56);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (124, 50, 'transfer', 'Transfer to savings', 'failed', 'EUR', 815.58);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (49, 132, 'transfer', 'Account transfer', 'completed', 'JPY', 3477.14);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (39, 154, 'transfer', 'Online transfer', 'completed', 'USD', 2367.07);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (22, 130, 'transfer', 'Internal transfer', 'completed', 'AUD', 751.73);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (147, 106, 'transfer', 'Account transfer', 'completed', 'USD', 3164.28);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (140, 154, 'transfer', 'Account transfer', 'failed', 'CAD', 4730.71);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (127, 131, 'transfer', 'Payment transfer', 'completed', 'USD', 926.86);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (99, 22, 'transfer', 'Internal transfer', 'completed', 'USD', 3339.2);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (113, 100, 'transfer', null, 'pending', 'GBP', 1305.9);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (43, 123, 'transfer', 'Transfer to savings', 'cancelled', 'CAD', 2284.86);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (43, 32, 'transfer', 'Account transfer', 'failed', 'CHF', 430.21);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (107, 12, 'transfer', 'Internal transfer', 'completed', 'USD', 1306.08);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (137, 43, 'transfer', 'Account transfer', 'pending', 'USD', 3562.69);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (36, 10, 'transfer', 'Internal transfer', 'cancelled', 'JPY', 3602.99);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (126, 42, 'transfer', 'Internal transfer', 'pending', 'JPY', 436.0);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (106, 28, 'transfer', 'Internal transfer', 'pending', 'USD', 2474.14);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (34, 45, 'transfer', 'Transfer to savings', 'cancelled', 'USD', 1186.43);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (115, 122, 'transfer', 'Transfer to savings', 'completed', 'MXN', 795.64);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (52, 35, 'transfer', 'Account transfer', 'pending', 'JPY', 4605.45);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (99, 23, 'transfer', 'Transfer to savings', 'completed', 'JPY', 3346.47);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (35, 99, 'transfer', 'Payment transfer', 'cancelled', 'MXN', 1135.47);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (141, 107, 'transfer', 'Account transfer', 'cancelled', 'USD', 4850.03);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (64, 126, 'transfer', 'Internal transfer', 'cancelled', 'AUD', 1557.99);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (114, 94, 'transfer', 'Account transfer', 'cancelled', 'CAD', 3502.72);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (72, 11, 'transfer', null, 'pending', 'EUR', 3426.51);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (105, 19, 'transfer', 'Account transfer', 'completed', 'CAD', 2683.7);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (129, 72, 'transfer', 'Online transfer', 'pending', 'EUR', 4151.32);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (151, 73, 'transfer', null, 'completed', 'CAD', 4820.91);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (79, 11, 'transfer', 'Internal transfer', 'completed', 'EUR', 2048.82);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (13, 51, 'transfer', 'Internal transfer', 'completed', 'USD', 3315.31);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (109, 8, 'transfer', 'Internal transfer', 'failed', 'CAD', 1237.75);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (110, 11, 'transfer', 'Payment transfer', 'failed', 'EUR', 4938.76);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (94, 70, 'transfer', 'Transfer to savings', 'completed', 'CHF', 2807.43);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (160, 79, 'transfer', 'Transfer to savings', 'cancelled', 'INR', 174.53);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (29, 39, 'transfer', 'Account transfer', 'completed', 'GBP', 4759.68);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (79, 75, 'transfer', 'Internal transfer', 'completed', 'INR', 381.07);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (143, 55, 'transfer', 'Transfer to savings', 'pending', 'CHF', 97.99);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (159, 126, 'transfer', 'Payment transfer', 'failed', 'AUD', 2236.22);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (121, 86, 'transfer', 'Account transfer', 'completed', 'INR', 2371.6);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (86, 78, 'transfer', 'Internal transfer', 'completed', 'GBP', 2831.8);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (26, 146, 'transfer', 'Internal transfer', 'completed', 'CAD', 3255.54);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (151, 51, 'transfer', 'Internal transfer', 'completed', 'EUR', 1009.99);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (30, 2, 'transfer', 'Payment transfer', 'failed', 'EUR', 3021.82);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (12, 120, 'transfer', 'Internal transfer', 'pending', 'AUD', 3928.4);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (128, 107, 'transfer', 'Account transfer', 'completed', 'EUR', 2837.4);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (64, 15, 'transfer', 'Online transfer', 'pending', 'GBP', 1354.89);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (15, 26, 'transfer', 'Payment transfer', 'completed', 'JPY', 4737.65);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (2, 88, 'transfer', 'Transfer to savings', 'completed', 'JPY', 892.07);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (71, 54, 'transfer', 'Account transfer', 'cancelled', 'USD', 4181.88);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (73, 58, 'transfer', 'Online transfer', 'completed', 'MXN', 1438.57);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (46, 35, 'transfer', 'Online transfer', 'completed', 'JPY', 4068.86);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (34, 126, 'transfer', 'Transfer to savings', 'pending', 'AUD', 1581.42);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (96, 100, 'transfer', 'Transfer to savings', 'failed', 'CAD', 4653.99);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (144, 160, 'transfer', 'Online transfer', 'cancelled', 'MXN', 3575.84);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (134, 43, 'transfer', null, 'pending', 'JPY', 3950.28);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (6, 143, 'transfer', 'Internal transfer', 'failed', 'CHF', 2526.86);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (92, 85, 'transfer', 'Transfer to savings', 'completed', 'CAD', 790.26);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (76, 108, 'transfer', 'Transfer to savings', 'failed', 'USD', 2529.44);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (115, 26, 'transfer', 'Transfer to savings', 'failed', 'CHF', 2490.17);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (109, 100, 'transfer', 'Payment transfer', 'pending', 'JPY', 579.34);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (62, 75, 'transfer', 'Transfer to savings', 'completed', 'USD', 2435.17);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (145, 20, 'transfer', 'Transfer to savings', 'completed', 'USD', 690.73);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (15, 68, 'transfer', 'Payment transfer', 'failed', 'AUD', 3705.31);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (25, 100, 'transfer', null, 'completed', 'AUD', 293.19);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (61, 105, 'transfer', 'Online transfer', 'completed', 'EUR', 1646.52);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (159, 157, 'transfer', 'Internal transfer', 'completed', 'USD', 1063.66);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (107, 102, 'transfer', 'Internal transfer', 'completed', 'USD', 2620.81);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (93, 21, 'transfer', 'Account transfer', 'failed', 'MXN', 761.38);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (48, 111, 'transfer', null, 'pending', 'AUD', 3812.06);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (13, 4, 'transfer', 'Internal transfer', 'completed', 'JPY', 4971.42);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (140, 105, 'transfer', 'Online transfer', 'completed', 'CHF', 2844.12);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (28, 102, 'transfer', 'Account transfer', 'failed', 'EUR', 4215.36);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (88, 99, 'transfer', 'Account transfer', 'completed', 'USD', 3752.66);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (10, 140, 'transfer', 'Account transfer', 'completed', 'INR', 3142.3);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (160, 25, 'transfer', 'Payment transfer', 'cancelled', 'AUD', 2715.83);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (160, 16, 'transfer', 'Account transfer', 'pending', 'JPY', 2067.82);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (127, 78, 'transfer', 'Internal transfer', 'cancelled', 'USD', 3316.97);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (88, 160, 'transfer', 'Account transfer', 'failed', 'EUR', 491.47);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (143, 133, 'transfer', 'Online transfer', 'cancelled', 'CHF', 3763.57);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (105, 78, 'transfer', 'Payment transfer', 'completed', 'CAD', 1460.34);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (103, 89, 'transfer', 'Online transfer', 'cancelled', 'AUD', 1868.31);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (3, 55, 'transfer', null, 'completed', 'CAD', 4913.41);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (38, 33, 'transfer', 'Transfer to savings', 'completed', 'AUD', 231.47);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (5, 5, 'transfer', 'Online transfer', 'completed', 'EUR', 4315.82);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (14, 13, 'transfer', null, 'completed', 'AUD', 1984.3);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (42, 32, 'transfer', 'Online transfer', 'completed', 'USD', 3780.66);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (125, 20, 'transfer', 'Online transfer', 'pending', 'GBP', 3241.04);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (90, 99, 'transfer', 'Internal transfer', 'failed', 'USD', 4598.67);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (45, 123, 'transfer', null, 'completed', 'JPY', 2359.66);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (25, 139, 'transfer', 'Internal transfer', 'completed', 'CHF', 2831.02);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (147, 37, 'transfer', 'Account transfer', 'cancelled', 'MXN', 1577.81);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (55, 105, 'transfer', 'Account transfer', 'pending', 'USD', 125.34);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (140, 149, 'transfer', 'Online transfer', 'completed', 'CHF', 2952.0);
insert into mock_transfers (to_account_id, from_account_id, transaction_type, transaction_description, transaction_status, currency, amount) values (98, 146, 'transfer', 'Online transfer', 'cancelled', 'USD', 1800.48);

INSERT INTO transactions (
  to_account_id,
  from_account_id,
  transaction_type,
  transaction_description,
  transaction_status,
  currency,
  amount
)
SELECT
  CASE
    WHEN to_account_id = from_account_id THEN
      CASE
        WHEN to_account_id = 1 THEN 2
        ELSE to_account_id - 1
      END
    ELSE to_account_id
  END AS to_account_id,
  from_account_id,
  transaction_type,
  transaction_description,
  transaction_status,
  currency,
  amount
FROM mock_transfers;