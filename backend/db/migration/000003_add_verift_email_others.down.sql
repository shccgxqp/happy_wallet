ALTER TABLE "verify_emails" DROP CONSTRAINT IF EXISTS verify_emails_user_id_fkey;

DROP TABLE IF EXISTS "verify_emails";

ALTER TABLE "users" DROP COLUMN IF EXISTS "is_email_verified";
ALTER TABLE "users" DROP COLUMN IF EXISTS "password_change_at";
ALTER TABLE "sessions" DROP COLUMN IF EXISTS "username";