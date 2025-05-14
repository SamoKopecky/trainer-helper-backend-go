ALTER TABLE week_day
ADD COLUMN timeslot_id integer;

CREATE INDEX IF NOT EXISTS idx_timesloit_id ON week_day USING btree (timeslot_id);

ALTER TABLE exercise
RENAME COLUMN timeslot_id TO week_day_id;
