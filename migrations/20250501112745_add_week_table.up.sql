CREATE TABLE IF NOT EXISTS week (
  id integer NOT NULL,
  user_id character varying NOT NULL,
  start_date timestamp without time zone NOT NULL,
  label integer NOT NULL,
  block_label integer NOT NULL,
  monday character varying,
  tuesday character varying,
  wednesday character varying,
  thursday character varying,
  friday character varying,
  saturday character varying,
  sunday character varying,
  created_at timestamp without time zone NOT NULL,
  updated_at timestamp without time zone NOT NULL
);

CREATE INDEX id_week_user_id ON week USING btree (user_id);

CREATE SEQUENCE week_id_seq AS integer START
WITH
  1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

ALTER TABLE ONLY week ALTER COLUMN id SET DEFAULT nextval('week_id_seq'::regclass);
