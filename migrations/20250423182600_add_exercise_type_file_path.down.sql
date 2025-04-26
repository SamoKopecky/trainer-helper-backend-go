ALTER TABLE exercise_type
RENAME COLUMN youtube_link TO media_address;

ALTER TABLE exercise_type
DROP COLUMN file_path;
