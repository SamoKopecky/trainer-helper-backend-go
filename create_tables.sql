--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2 (Debian 17.2-1.pgdg120+1)
-- Dumped by pg_dump version 17.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: exercise; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.exercise (
    id integer NOT NULL,
    timeslot_id integer NOT NULL,
    group_id integer NOT NULL,
    set_type character varying NOT NULL,
    note character varying,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.exercise OWNER TO root;

--
-- Name: exercise_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.exercise_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.exercise_id_seq OWNER TO root;

--
-- Name: exercise_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.exercise_id_seq OWNED BY public.exercise.id;


--
-- Name: person; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.person (
    id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.person OWNER TO root;

--
-- Name: person_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.person_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.person_id_seq OWNER TO root;

--
-- Name: person_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.person_id_seq OWNED BY public.person.id;


--
-- Name: seaql_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.seaql_migrations (
    version character varying NOT NULL,
    applied_at bigint NOT NULL
);


ALTER TABLE public.seaql_migrations OWNER TO root;

--
-- Name: timeslot; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.timeslot (
    id integer NOT NULL,
    trainer_id integer NOT NULL,
    name character varying NOT NULL,
    start timestamp without time zone NOT NULL,
    "end" timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    user_id integer
);


ALTER TABLE public.timeslot OWNER TO root;

--
-- Name: timeslot_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.timeslot_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.timeslot_id_seq OWNER TO root;

--
-- Name: timeslot_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.timeslot_id_seq OWNED BY public.timeslot.id;


--
-- Name: work_set; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.work_set (
    id integer NOT NULL,
    reps integer NOT NULL,
    intensity character varying NOT NULL,
    rpe integer,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    exercise_id integer NOT NULL
);


ALTER TABLE public.work_set OWNER TO root;

--
-- Name: work_set_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.work_set_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_set_id_seq OWNER TO root;

--
-- Name: work_set_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.work_set_id_seq OWNED BY public.work_set.id;


--
-- Name: exercise id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.exercise ALTER COLUMN id SET DEFAULT nextval('public.exercise_id_seq'::regclass);


--
-- Name: person id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.person ALTER COLUMN id SET DEFAULT nextval('public.person_id_seq'::regclass);


--
-- Name: timeslot id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.timeslot ALTER COLUMN id SET DEFAULT nextval('public.timeslot_id_seq'::regclass);


--
-- Name: work_set id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.work_set ALTER COLUMN id SET DEFAULT nextval('public.work_set_id_seq'::regclass);


--
-- Name: exercise exercise_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.exercise
    ADD CONSTRAINT exercise_pkey PRIMARY KEY (id);


--
-- Name: person person_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.person
    ADD CONSTRAINT person_pkey PRIMARY KEY (id);


--
-- Name: seaql_migrations seaql_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.seaql_migrations
    ADD CONSTRAINT seaql_migrations_pkey PRIMARY KEY (version);


--
-- Name: timeslot timeslot_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.timeslot
    ADD CONSTRAINT timeslot_pkey PRIMARY KEY (id);


--
-- Name: work_set work_set_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.work_set
    ADD CONSTRAINT work_set_pkey PRIMARY KEY (id);


--
-- Name: idx_exercise_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_exercise_id ON public.work_set USING btree (exercise_id);


--
-- Name: idx_exercise_timeslot_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_exercise_timeslot_id ON public.exercise USING btree (timeslot_id);


--
-- Name: idx_group_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_group_id ON public.exercise USING btree (group_id);


--
-- Name: idx_start; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_start ON public.timeslot USING btree (start);


--
-- Name: idx_trainer_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_trainer_id ON public.timeslot USING btree (trainer_id);


--
-- Name: idx_user_id; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_user_id ON public.timeslot USING btree (user_id);


--
-- PostgreSQL database dump complete
--

