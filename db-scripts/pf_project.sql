-- Table: public.pf_project

-- DROP TABLE IF EXISTS public.pf_project;

CREATE TABLE IF NOT EXISTS public.pf_project
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    title text COLLATE pg_catalog."default" NOT NULL,
    owner_id uuid,
    description text COLLATE pg_catalog."default",
    is_archived boolean DEFAULT false,
    CONSTRAINT pf_project_pkey PRIMARY KEY (id),
    CONSTRAINT fk_pf_project_pf_user FOREIGN KEY (owner_id)
        REFERENCES public.pf_user (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_project
    OWNER to postgres;
-- Index: idx_pf_project_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_project_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_project_deleted_at
    ON public.pf_project USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: idx_pf_project_title

-- DROP INDEX IF EXISTS public.idx_pf_project_title;

CREATE INDEX IF NOT EXISTS idx_pf_project_title
    ON public.pf_project USING btree
    (title COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;