ALTER TABLE exercise_type
RENAME COLUMN media_address TO youtube_link;

ALTER TABLE exercise_type
ADD COLUMN file_path varchar;
