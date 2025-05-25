ALTER TABLE week_day
DROP COLUMN timeslot_id;

DROP INDEX idx_timeslot_id;

ALTER TABLE exercise
RENAME COLUMN week_day_id TO timeslot_id;
