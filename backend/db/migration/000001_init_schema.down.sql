ALTER TABLE "teams" DROP CONSTRAINT IF EXISTS teams_owner_fkey;
ALTER TABLE "team_members" DROP CONSTRAINT IF EXISTS team_members_team_id_fkey;
ALTER TABLE "team_members" DROP CONSTRAINT IF EXISTS team_members_linked_user_id_fkey;
ALTER TABLE "expenses" DROP CONSTRAINT IF EXISTS expenses_team_id_fkey;
ALTER TABLE "expense_details" DROP CONSTRAINT IF EXISTS expense_details_member_id_fkey;
ALTER TABLE "expense_details" DROP CONSTRAINT IF EXISTS expense_details_expense_id_fkey;

DROP TABLE IF EXISTS "expense_details";
DROP TABLE IF EXISTS "expenses";
DROP TABLE IF EXISTS "team_members";
DROP TABLE IF EXISTS "teams";
DROP TABLE IF EXISTS "users";