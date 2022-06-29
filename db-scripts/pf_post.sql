-- Table: public.pf_post

-- DROP TABLE IF EXISTS public.pf_post;

CREATE SEQUENCE public.pf_post_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.pf_post_id_seq
    OWNER TO postgres;

CREATE TABLE IF NOT EXISTS public.pf_post
(
    id bigint NOT NULL DEFAULT nextval('pf_post_id_seq'::regclass),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    title text NOT NULL COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    votes bigint DEFAULT 0,
    status text COLLATE pg_catalog."default",
    pf_board_id uuid,
    pf_category_id bigint,
    CONSTRAINT pf_post_pkey PRIMARY KEY (id),
    CONSTRAINT fk_pf_board_ticket FOREIGN KEY (pf_board_id)
        REFERENCES public.pf_board (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_pf_post_pf_category FOREIGN KEY (pf_category_id)
        REFERENCES public.pf_category (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_post
    OWNER to postgres;
-- Index: idx_pf_board_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_tickets_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_ticket_deleted_at
    ON public.pf_post USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: idx_pf_post_status

-- DROP INDEX IF EXISTS public.idx_pf_post_status;

CREATE INDEX IF NOT EXISTS idx_pf_post_status
    ON public.pf_post USING btree
    (status COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;