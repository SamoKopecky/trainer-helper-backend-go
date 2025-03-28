CREATE TABLE exercise (
    id integer NOT NULL,
    timeslot_id integer NOT NULL,
    group_id integer NOT NULL,
    set_type character varying NOT NULL,
    note character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE TABLE timeslot (
    id integer NOT NULL,
    trainer_id character varying NOT NULL,
    trainee_id character varying,
    name character varying NOT NULL,
    start timestamp without time zone NOT NULL,
    "end" timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone 
);

CREATE TABLE work_set (
    id integer NOT NULL,
    exercise_id integer NOT NULL,
    reps integer NOT NULL,
    intensity character varying NOT NULL,
    rpe integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


CREATE SEQUENCE work_set_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
CREATE SEQUENCE exercise_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
CREATE SEQUENCE timeslot_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE ONLY exercise ALTER COLUMN id SET DEFAULT nextval('exercise_id_seq'::regclass);
ALTER TABLE ONLY timeslot ALTER COLUMN id SET DEFAULT nextval('timeslot_id_seq'::regclass);
ALTER TABLE ONLY work_set ALTER COLUMN id SET DEFAULT nextval('work_set_id_seq'::regclass);


ALTER TABLE ONLY exercise ADD CONSTRAINT exercise_pkey PRIMARY KEY (id);
ALTER TABLE ONLY timeslot ADD CONSTRAINT timeslot_pkey PRIMARY KEY (id);
ALTER TABLE ONLY work_set ADD CONSTRAINT work_set_pkey PRIMARY KEY (id);


CREATE INDEX idx_exercise_id ON work_set USING btree (exercise_id);
CREATE INDEX idx_timeslot_id ON exercise USING btree (timeslot_id);
CREATE INDEX idx_group_id ON exercise USING btree (group_id);
CREATE INDEX idx_start ON timeslot USING btree (start);
CREATE INDEX idx_trainer_id ON timeslot USING btree (trainer_id);
CREATE INDEX idx_trainee_id ON timeslot USING btree (trainee_id);

