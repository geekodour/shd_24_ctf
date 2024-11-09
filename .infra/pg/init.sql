-- How to use this?
-- - This is supposed to run manually when you first setup postgres
-- - This is supposed to run before you apply any migration or apply any kind of
--   DDL/DML to the database
-- - This script has statements which will error out if run more than once
-- - You don't really need to run this script on existing db, it'll already have
--   those changes
-- - This is NOT a migration script, this is pure DB maintenance script. The
--   database which are being created here are used to connect to the db in
--   which we run migrations etc.
-- - We're not running this script as part of postgres init because then it'll
--   run everytime postgres starts, which is something that we don't need.
-- - Yes, if you make such operations directly on the DB, make sure to record it
--   here so that there's a record of it in this file. This "is" a manual thing
--   you need to take care of.
-- - You'll only need this when trying to create a fresh replica of the database
--   which would not be the case most of the time, if you make changes in the
--   database, they'll stay in the database and you'll keep mutating it etc. and
--   then you'll have backups too etc. In any other cases, when you need a dump
--   of the db, pg_dumpall --schema-only can help with it so you probably won't
--   need this file. But this is still a nice but manual way to keep track of
--   what's going on the db i think.
-- - There are terraform providers to do sort of the things we're doing here but
--   we're not going to use it, cuz why?
--   - https://registry.terraform.io/providers/cyrilgdn/postgresql/latest/docs
--
-- NOTE:
-- When creating new ROLE, they'll be able to connect to other db and see basic info
-- This is fine because we don't expect people to be knowing the password, we
-- can be more strict about limiting access to certain databases we create.
-- However, postgres by design gives some access to any new role created.
-- See following for more info:
-- https://stackoverflow.com/questions/70003000/postgresql-restricting-postgresql-user-access-to-a-particular-database
-- https://stackoverflow.com/questions/6884020/why-new-user-in-postgresql-can-connect-to-all-databases
--
-- NOTE: To generate SCRAM-SHA verifier you can use the ./encrypt_password.py script
-- - It's OK for this thing to be publicly accessible, SCRAM is great!
--
-- How to run this script? (You need to have psql installed)
-- psql -h <hostname> -U postgres -d postgres -f ./init_pg.sql

\c postgres postgres
CREATE ROLE app_user WITH LOGIN PASSWORD 'SCRAM-SHA-256$4096:+fI477HM+2+a46V/vhQQsA==$sXKIn8DBV23aUOVuJA8jA76mb0wZcXk00GQtqgUD33A=:6ALmD1Wt2BvkZvkMPrdxyBSlkTvEtHgnIdiHjr5+0R8=';
CREATE DATABASE boilerpress WITH OWNER app_user;
\c boilerpress
-- NOTE: Make sure you re-start the postgres instance(if running) after installing the extension.
CREATE EXTENSION "pg_uuidv7"; -- (extension installed via nixpkgs)
-- CREATE EXTENSION "uuidv7"; -- (extension installed via others)
