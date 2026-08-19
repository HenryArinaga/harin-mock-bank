CREATE VIEW customer_balances_json AS
SELECT
  customer_id,
  first_name,
  last_name,
  email,
  jsonb_object_agg(currency, total_balance ORDER BY currency) AS balances_by_currency
FROM customer_currency_balances
GROUP BY
  customer_id,
  first_name,
  last_name,
  email;