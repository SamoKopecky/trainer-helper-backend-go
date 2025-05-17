ALTER TABLE week_day
ADD COLUMN deleted_at timestamp without time zone;

ALTER TABLE timeslot
DROP COLUMN name;
