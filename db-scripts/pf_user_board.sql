-- Table: public.pf_user_board

-- DROP TABLE IF EXISTS public.pf_user_board;

CREATE TABLE IF NOT EXISTS public.pf_user_board
(
    pf_board_id bigint NOT NULL,
    pf_user_id bigint NOT NULL,
    CONSTRAINT pf_user_board_pkey PRIMARY KEY (pf_board_id, pf_user_id),
    CONSTRAINT fk_pf_user_board_pf_board FOREIGN KEY (pf_board_id)
        REFERENCES public.pf_board (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_pf_user_board_pf_user FOREIGN KEY (pf_user_id)
        REFERENCES public.pf_user (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.pf_user_board
    OWNER to postgres;