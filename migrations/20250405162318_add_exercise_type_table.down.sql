DROP TABLE IF EXISTS exercise_type;

DROP SEQUENCE IF EXISTS exercise_type_id_seq;

-- TODO: Check if this works and revert nullable
UPDATE exercise
SET
  exercise_type = null
ALTER TABLE exercise
ALTER COLUMN exercise_type_id TYPE character varying;

ALTER TABLE exercise
RENAME COLUMN exercise_type_id TO set_type;
