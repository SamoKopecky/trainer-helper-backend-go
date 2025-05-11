-- =================================================================
-- Block Table
-- =================================================================

CREATE SEQUENCE IF NOT EXISTS block_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS block (
    id integer NOT NULL DEFAULT nextval('block_id_seq'::regclass),
    user_id character varying NOT NULL,
    label integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
	deleted_at timestamp without time zone
);

ALTER TABLE ONLY block
    ADD CONSTRAINT block_pkey PRIMARY KEY (id);

CREATE INDEX IF NOT EXISTS idx_block_user_id ON block USING btree (user_id);


-- =================================================================
-- Week Table (Replaces original structure)
-- =================================================================

CREATE SEQUENCE IF NOT EXISTS week_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS week (
    id integer NOT NULL DEFAULT nextval('week_id_seq'::regclass),
    user_id character varying NOT NULL,
    block_id integer NOT NULL,
    start_date date NOT NULL,
    label integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
	deleted_at timestamp without time zone
);

ALTER TABLE ONLY week
    ADD CONSTRAINT week_pkey PRIMARY KEY (id);

CREATE INDEX IF NOT EXISTS idx_week_user_id ON week USING btree (user_id);

CREATE INDEX IF NOT EXISTS idx_week_block_id ON week USING btree (block_id);


-- =================================================================
-- WeekDay Table
-- =================================================================

CREATE SEQUENCE IF NOT EXISTS week_day_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS week_day (
    id integer NOT NULL DEFAULT nextval('week_day_id_seq'::regclass),
    user_id character varying NOT NULL,
    week_id integer NOT NULL,
    day_date date NOT NULL,
    name character varying NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

ALTER TABLE ONLY week_day
    ADD CONSTRAINT week_day_pkey PRIMARY KEY (id);


CREATE INDEX IF NOT EXISTS idx_week_day_week_id ON week_day USING btree (week_id);
