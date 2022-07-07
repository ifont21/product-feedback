-- Table: public.pf_board

-- DROP TABLE IF EXISTS public.pf_board;

CREATE TABLE IF NOT EXISTS public.pf_board
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    pf_project_id uuid,
    CONSTRAINT pf_board_pkey PRIMARY KEY (id),
    CONSTRAINT fk_pf_project_pf_boards FOREIGN KEY (pf_project_id)
        REFERENCES public.pf_project (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_board
    OWNER to postgres;
-- Index: idx_pf_board_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_board_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_board_deleted_at
    ON public.pf_board USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;
-- Index: idx_pf_boards_name

-- DROP INDEX IF EXISTS public.idx_pf_boards_name;

CREATE INDEX IF NOT EXISTS idx_pf_boards_name
    ON public.pf_board USING btree
    (name COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;