-- Table: public.pf_user_project

-- DROP TABLE IF EXISTS public.pf_user_project;

CREATE TABLE IF NOT EXISTS public.pf_user_project
(
    pf_project_id uuid NOT NULL,
    pf_user_id uuid NOT NULL,
    CONSTRAINT pf_user_project_pkey PRIMARY KEY (pf_project_id, pf_user_id),
    CONSTRAINT fk_pf_user_project_pf_project FOREIGN KEY (pf_project_id)
        REFERENCES public.pf_project (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_pf_user_project_pf_user FOREIGN KEY (pf_user_id)
        REFERENCES public.pf_user (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_user_project
    OWNER to postgres;