--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: oauth_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oauth_users (
    oauth_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.oauth_users OWNER TO postgres;

--
-- Name: oauths; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oauths (
    id integer NOT NULL,
    provider character varying NOT NULL,
    provider_id character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.oauths OWNER TO postgres;

--
-- Name: oauths_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.oauths_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.oauths_id_seq OWNER TO postgres;

--
-- Name: oauths_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.oauths_id_seq OWNED BY public.oauths.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    blocked boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: zgroup_members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.zgroup_members (
    zgroup_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.zgroup_members OWNER TO postgres;

--
-- Name: zgroups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.zgroups (
    id integer NOT NULL,
    name character varying NOT NULL,
    archived boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.zgroups OWNER TO postgres;

--
-- Name: zgroups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.zgroups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.zgroups_id_seq OWNER TO postgres;

--
-- Name: zgroups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.zgroups_id_seq OWNED BY public.zgroups.id;


--
-- Name: oauths id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauths ALTER COLUMN id SET DEFAULT nextval('public.oauths_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: zgroups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zgroups ALTER COLUMN id SET DEFAULT nextval('public.zgroups_id_seq'::regclass);


--
-- Name: oauth_users oauth_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauth_users
    ADD CONSTRAINT oauth_users_pkey PRIMARY KEY (oauth_id, user_id);


--
-- Name: oauths oauths_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauths
    ADD CONSTRAINT oauths_pkey PRIMARY KEY (id);


--
-- Name: oauths oauths_provider_provider_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauths
    ADD CONSTRAINT oauths_provider_provider_id_key UNIQUE (provider, provider_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: zgroup_members zgroup_members_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zgroup_members
    ADD CONSTRAINT zgroup_members_pkey PRIMARY KEY (zgroup_id, user_id);


--
-- Name: zgroups zgroups_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zgroups
    ADD CONSTRAINT zgroups_name_key UNIQUE (name);


--
-- Name: zgroups zgroups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zgroups
    ADD CONSTRAINT zgroups_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: oauth_users fk_oauth; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauth_users
    ADD CONSTRAINT fk_oauth FOREIGN KEY (oauth_id) REFERENCES public.oauths(id);


--
-- Name: oauth_users fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oauth_users
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: zgroup_members fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zgroup_members
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: zgroup_members fk_zgroup; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.zgroup_members
    ADD CONSTRAINT fk_zgroup FOREIGN KEY (zgroup_id) REFERENCES public.zgroups(id);


--
-- PostgreSQL database dump complete
--

