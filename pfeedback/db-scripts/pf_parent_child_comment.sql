-- Table: public.pf_parent_child_comment

-- DROP TABLE IF EXISTS public.pf_parent_child_comment;

CREATE TABLE IF NOT EXISTS public.pf_parent_child_comment
(
    id bigint NOT NULL DEFAULT nextval('pf_parent_child_comment_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    pf_parent_comment_id bigint,
    pf_child_comment_id bigint,
    CONSTRAINT pf_parent_child_comment_pkey PRIMARY KEY (id),
    CONSTRAINT fk_pf_parent_child_comment_pf_child_comment FOREIGN KEY (pf_child_comment_id)
        REFERENCES public.pf_comment (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_pf_parent_child_comment_pf_parent_comment FOREIGN KEY (pf_parent_comment_id)
        REFERENCES public.pf_comment (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_parent_child_comment
    OWNER to postgres;
-- Index: idx_pf_parent_child_comments_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_parent_child_comments_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_parent_child_comment_deleted_at
    ON public.pf_parent_child_comment USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;