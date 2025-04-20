CREATE TABLE exercise_type (
  id integer NOT NULL,
  user_id character varying NOT NULL,
  name character varying NOT NULL,
  note character varying,
  media_type character varying,
  media_address character varying,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE SEQUENCE exercise_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY exercise_type ALTER COLUMN id SET DEFAULT nextval('exercise_type_id_seq'::regclass);

ALTER TABLE ONLY exercise_type ADD CONSTRAINT exercise_type_key PRIMARY KEY (id);

ALTER TABLE exercise RENAME COLUMN set_type TO exercise_type_id;

ALTER TABLE exercise ALTER COLUMN exercise_type_id TYPE integer USING (NULLIF(exercise_type_id, '')::integer);
ALTER TABLE exercise ALTER COLUMN exercise_type_id DROP NOT NULL;
