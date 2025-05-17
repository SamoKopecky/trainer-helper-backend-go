ALTER TABLE week_day
DROP COLUMN deleted_at;

ALTER TABLE timeslot
ADD COLUMN name character varying NOT NULL;
