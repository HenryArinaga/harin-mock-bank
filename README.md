# Harin Mock Bank

Local mock banking database using PostgreSQL.

## Requirements

- PostgreSQL 16+
- `psql`
- Local database named `bank_app`

## Environment

Create `.env` in the project root:

```env
DATABASE_URL=postgresql://YOUR_DB_USER:YOUR_DB_PASSWORD@localhost:5432/bank_app