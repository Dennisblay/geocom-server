-- Drop foreign key constraints
ALTER TABLE "credit" DROP CONSTRAINT IF EXISTS "credit_user_id_fkey";
ALTER TABLE "computation" DROP CONSTRAINT IF EXISTS "computation_user_id_fkey";
ALTER TABLE "ussd_transaction" DROP CONSTRAINT IF EXISTS "ussd_transaction_user_id_fkey";
ALTER TABLE "token" DROP CONSTRAINT IF EXISTS "token_user_id_fkey";

-- Drop indices
DROP INDEX IF EXISTS "credit_user_id_idx";
DROP INDEX IF EXISTS "computation_computation_type_idx";
DROP INDEX IF EXISTS "ussd_transaction_transaction_id_idx";
DROP INDEX IF EXISTS "token_token_idx";
DROP INDEX IF EXISTS "token_user_id_idx";
DROP INDEX IF EXISTS "user_phone_idx";
DROP INDEX IF EXISTS "user_email_idx";

-- Drop tables
DROP TABLE IF EXISTS "credit";
DROP TABLE IF EXISTS "computation";
DROP TABLE IF EXISTS "ussd_transaction";
DROP TABLE IF EXISTS "token";
DROP TABLE IF EXISTS "user";
