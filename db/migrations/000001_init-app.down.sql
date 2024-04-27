DROP EXTENSION IF EXISTS "uuid-ossp";

-- Step 2: Drop the users table
DROP TABLE IF EXISTS sending_emails;
DROP TABLE IF EXISTS subscribers;
DROP TABLE IF EXISTS newsletters;

DROP TYPE IF EXISTS status;