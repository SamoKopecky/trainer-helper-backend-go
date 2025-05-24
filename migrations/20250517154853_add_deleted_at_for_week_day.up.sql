ALTER TABLE week_day
ADD COLUMN deleted_at timestamp without time zone;

ALTER TABLE timeslot
DROP COLUMN name;

CREATE EXTENSION IF NOT EXISTS btree_gist;

CREATE INDEX idx_start_date_gist_week ON week USING GIST (start_date);

CREATE INDEX idx_week_day_week_day ON week_day USING btree (day_date);
