ALTER TABLE week_day
DROP COLUMN deleted_at;

ALTER TABLE timeslot
ADD COLUMN name character varying NOT NULL;

DROP INDEX IF EXISTS idx_start_date_gist_week;

DROP INDEX IF EXISTS idx_week_day_week_day;
