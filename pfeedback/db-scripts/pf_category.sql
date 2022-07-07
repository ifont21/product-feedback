-- Table: public.pf_category

-- DROP TABLE IF EXISTS public.pf_category;

CREATE SEQUENCE public.pf_category_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.pf_category_id_seq
    OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.pf_category
(
    id bigint NOT NULL DEFAULT nextval('pf_category_id_seq'::regclass),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    CONSTRAINT pf_category_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_category
    OWNER to postgres;
-- Index: idx_pf_categories_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_categories_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_category_deleted_at
    ON public.pf_category USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;