-- Table: public.pf_comment

-- DROP TABLE IF EXISTS public.pf_comment;

CREATE TABLE IF NOT EXISTS public.pf_comment
(
    id bigint NOT NULL DEFAULT nextval('pf_comment_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    comment text COLLATE pg_catalog."default",
    pf_user_id bigint,
    pf_ticket_id bigint,
    CONSTRAINT pf_comment_pkey PRIMARY KEY (id),
    CONSTRAINT fk_pf_ticket_pf_comment FOREIGN KEY (pf_ticket_id)
        REFERENCES public.pf_ticket (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_pf_user_comment FOREIGN KEY (pf_user_id)
        REFERENCES public.pf_user (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_comment
    OWNER to postgres;
-- Index: idx_pf_comments_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_comments_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_comment_deleted_at
    ON public.pf_comment USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;