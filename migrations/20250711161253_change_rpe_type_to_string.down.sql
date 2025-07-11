-- Step 1: Update non-integer values to a default integer (e.g., 0)
UPDATE work_set
SET rpe = '0' -- Set to string '0' if rpe is text/varchar, it will be cast to integer 0 later
WHERE rpe IS NOT NULL AND rpe !~ '^-?[0-9]+$';

-- Step 2: Now run your ALTER TABLE statement
ALTER TABLE work_set
ALTER COLUMN rpe TYPE integer USING rpe::integer,
ALTER COLUMN rpe DROP NOT NULL;
