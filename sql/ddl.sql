CREATE TABLE public.client (
	id int4 NOT NULL,
	"limit" int4 NOT NULL,
	balance int4 NOT NULL,
	CONSTRAINT client_pk PRIMARY KEY (id)
);

CREATE TABLE public."transaction" (
	id serial4 NOT NULL,
	client_id int4 NOT NULL,
	"type" int4 NOT NULL,
	value int4 NOT NULL,
	description varchar NOT NULL, -- TODO: limit in 10 char
	created_at timestamp with time zone NOT NULL,
	CONSTRAINT transaction_pkey PRIMARY KEY (id)
);

ALTER TABLE public."transaction" ADD CONSTRAINT transaction_fk FOREIGN KEY (client_id) REFERENCES public.client(id);

CREATE INDEX transaction_client_id_idx ON public."transaction" (client_id);
