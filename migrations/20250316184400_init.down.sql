DROP INDEX IF EXISTS idx_exercise_id;

DROP INDEX IF EXISTS idx_timeslot_id;

DROP INDEX IF EXISTS idx_group_id;

DROP INDEX IF EXISTS idx_start;

DROP INDEX IF EXISTS idx_trainer_id;

DROP INDEX IF EXISTS idx_user_id;

DROP TABLE IF EXISTS work_set;

DROP TABLE IF EXISTS exercise;

DROP TABLE IF EXISTS timeslot;

DROP SEQUENCE IF EXISTS work_set_id_seq;

DROP SEQUENCE IF EXISTS exercise_id_seq;

DROP SEQUENCE IF EXISTS person_id_seq;

DROP SEQUENCE IF EXISTS timeslot_id_seq;
