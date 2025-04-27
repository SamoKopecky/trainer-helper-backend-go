ALTER TABLE work_set
ADD COLUMN deleted_at timestamp without time zone;

ALTER TABLE exercise
ADD COLUMN deleted_at timestamp without time zone;

ALTER TABLE exercise_type
ADD COLUMN deleted_at timestamp without time zone;
