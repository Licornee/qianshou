--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Homebrew)
-- Dumped by pg_dump version 14.5 (Homebrew)

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
-- Name: qianshou; Type: SCHEMA; Schema: -; Owner: yindongpeng
--

CREATE SCHEMA qianshou;


ALTER SCHEMA qianshou OWNER TO yindongpeng;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: relationship; Type: TABLE; Schema: qianshou; Owner: yindongpeng
--

CREATE TABLE qianshou.relationship (
    user_id bigint NOT NULL,
    other_user_id bigint NOT NULL,
    state character varying(10) NOT NULL,
    type character varying(15) NOT NULL
);


ALTER TABLE qianshou.relationship OWNER TO yindongpeng;

--
-- Name: relationship_other_user_id_seq; Type: SEQUENCE; Schema: qianshou; Owner: yindongpeng
--

CREATE SEQUENCE qianshou.relationship_other_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE qianshou.relationship_other_user_id_seq OWNER TO yindongpeng;

--
-- Name: relationship_other_user_id_seq; Type: SEQUENCE OWNED BY; Schema: qianshou; Owner: yindongpeng
--

ALTER SEQUENCE qianshou.relationship_other_user_id_seq OWNED BY qianshou.relationship.other_user_id;


--
-- Name: relationship_user_id_seq; Type: SEQUENCE; Schema: qianshou; Owner: yindongpeng
--

CREATE SEQUENCE qianshou.relationship_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE qianshou.relationship_user_id_seq OWNER TO yindongpeng;

--
-- Name: relationship_user_id_seq; Type: SEQUENCE OWNED BY; Schema: qianshou; Owner: yindongpeng
--

ALTER SEQUENCE qianshou.relationship_user_id_seq OWNED BY qianshou.relationship.user_id;


--
-- Name: user; Type: TABLE; Schema: qianshou; Owner: yindongpeng
--

CREATE TABLE qianshou."user" (
    id bigint NOT NULL,
    name character varying(20) NOT NULL,
    type character varying(10) NOT NULL
);


ALTER TABLE qianshou."user" OWNER TO yindongpeng;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: qianshou; Owner: yindongpeng
--

CREATE SEQUENCE qianshou.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE qianshou.user_id_seq OWNER TO yindongpeng;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: qianshou; Owner: yindongpeng
--

ALTER SEQUENCE qianshou.user_id_seq OWNED BY qianshou."user".id;


--
-- Name: relationship user_id; Type: DEFAULT; Schema: qianshou; Owner: yindongpeng
--

ALTER TABLE ONLY qianshou.relationship ALTER COLUMN user_id SET DEFAULT nextval('qianshou.relationship_user_id_seq'::regclass);


--
-- Name: relationship other_user_id; Type: DEFAULT; Schema: qianshou; Owner: yindongpeng
--

ALTER TABLE ONLY qianshou.relationship ALTER COLUMN other_user_id SET DEFAULT nextval('qianshou.relationship_other_user_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: qianshou; Owner: yindongpeng
--

ALTER TABLE ONLY qianshou."user" ALTER COLUMN id SET DEFAULT nextval('qianshou.user_id_seq'::regclass);


--
-- Data for Name: relationship; Type: TABLE DATA; Schema: qianshou; Owner: yindongpeng
--

COPY qianshou.relationship (user_id, other_user_id, state, type) FROM stdin;
11231244213	222333444	liked	relationship
11231244213	333222444	matched	relationship
11231244213	444333222	disliked	relationship
11231244213	21341231231	liked	relationship
21341231231	11231244213	matched	relationship
21341231231	11231244213	disliked	relationship
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: qianshou; Owner: yindongpeng
--

COPY qianshou."user" (id, name, type) FROM stdin;
21341231231	Bob	user
31231242322	Samantha	user
11231244213	test	user
222333444	henry	user
333222444	Tom	user
444333222	John	user
\.


--
-- Name: relationship_other_user_id_seq; Type: SEQUENCE SET; Schema: qianshou; Owner: yindongpeng
--

SELECT pg_catalog.setval('qianshou.relationship_other_user_id_seq', 1, false);


--
-- Name: relationship_user_id_seq; Type: SEQUENCE SET; Schema: qianshou; Owner: yindongpeng
--

SELECT pg_catalog.setval('qianshou.relationship_user_id_seq', 1, false);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: qianshou; Owner: yindongpeng
--

SELECT pg_catalog.setval('qianshou.user_id_seq', 5, true);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: qianshou; Owner: yindongpeng
--

ALTER TABLE ONLY qianshou."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: relationship relationship_other_user_id_fkey; Type: FK CONSTRAINT; Schema: qianshou; Owner: yindongpeng
--

ALTER TABLE ONLY qianshou.relationship
    ADD CONSTRAINT relationship_other_user_id_fkey FOREIGN KEY (other_user_id) REFERENCES qianshou."user"(id);


--
-- Name: relationship relationship_user_id_fkey; Type: FK CONSTRAINT; Schema: qianshou; Owner: yindongpeng
--

ALTER TABLE ONLY qianshou.relationship
    ADD CONSTRAINT relationship_user_id_fkey FOREIGN KEY (user_id) REFERENCES qianshou."user"(id);


--
-- PostgreSQL database dump complete
--

