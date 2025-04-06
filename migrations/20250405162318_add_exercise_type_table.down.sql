DROP TABLE IF EXISTS exercise_type;

DROP SEQUENCE IF EXISTS exercise_type_id_seq;

ALTER TABLE exercise
RENAME COLUMN exercise_type TO set_type;
