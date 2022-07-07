-- Table: public.pf_user

-- DROP TABLE IF EXISTS public.pf_user;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.pf_user
(
    id uuid DEFAULT uuid_generate_v4 (),
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone,
    first_name text COLLATE pg_catalog."default" NOT NULL,
    last_name text COLLATE pg_catalog."default" NOT NULL,
    username text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT pf_user_pkey PRIMARY KEY (id),
    CONSTRAINT pf_user_username_key UNIQUE (username)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_user
    OWNER to postgres;
-- Index: idx_pf_users_deleted_at

-- DROP INDEX IF EXISTS public.idx_pf_users_deleted_at;

CREATE INDEX IF NOT EXISTS idx_pf_user_deleted_at
    ON public.pf_user USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;