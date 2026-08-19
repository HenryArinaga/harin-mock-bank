#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

set -a
. ./.env
set +a

psql "$DATABASE_URL" -f db/reset.sql

psql "$DATABASE_URL" -f db/migrations/001_init.sql
psql "$DATABASE_URL" -f db/migrations/002_account_balances_view.sql
psql "$DATABASE_URL" -f db/migrations/003_transaction_history_view.sql
psql "$DATABASE_URL" -f db/migrations/004_customer_summary_view.sql
psql "$DATABASE_URL" -f db/migrations/005_customer_currency_balances_view.sql
psql "$DATABASE_URL" -f db/migrations/006_customer_balances_json_view.sql

psql "$DATABASE_URL" -f db/seeds/001_users.sql
psql "$DATABASE_URL" -f db/seeds/002_customers.sql
psql "$DATABASE_URL" -f db/seeds/003_accounts.sql
psql "$DATABASE_URL" -f db/seeds/004_deposits.sql
psql "$DATABASE_URL" -f db/seeds/005_withdrawals.sql
psql "$DATABASE_URL" -f db/seeds/006_transfers.sql
psql "$DATABASE_URL" -f db/seeds/007_ledger_entries.sql

psql "$DATABASE_URL" -c "select user_role, count(*) from users group by user_role order by user_role;"
psql "$DATABASE_URL" -c "select transaction_type, count(*) from transactions group by transaction_type order by transaction_type;"
psql "$DATABASE_URL" -c "select direction, count(*) from ledger_entries group by direction order by direction;"