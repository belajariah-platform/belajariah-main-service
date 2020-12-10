CREATE TABLE private.users (
	id serial NOT NULL,
	code varchar(20) NOT NULL DEFAULT public.generate_code('USC'::character varying, 'users'::character varying),
	role_code varchar(20) NOT NULL,
    email varchar (100) NOT NULL,
    username varchar(100) NULL,
    password varchar (225) NOT NULL,
    is_verified bool,
    is_active bool NOT NULL DEFAULT true,
	created_by varchar(100) NOT NULL,
	created_date timestamptz NOT NULL DEFAULT Now(),
	modified_by varchar(100) NULL,
	modified_date timestamptz NULL DEFAULT Now(),
	deleted_by varchar(100) NULL,
	deleted_date timestamptz NULL DEFAULT '1753-07-01 00:00:00+07:07:12'::timestamp with time zone,
	CONSTRAINT users_pk PRIMARY KEY (id),
	CONSTRAINT users_un UNIQUE (code),
	CONSTRAINT users_un_del UNIQUE (code, deleted_date),
    CONSTRAINT users_fk FOREIGN KEY (role_code) REFERENCES private.user_role(code)
);

CREATE TABLE private.user_role (
	id serial NOT NULL,
	code varchar(20) NOT NULL DEFAULT public.generate_code('URC'::character varying, 'user_role'::character varying),
	role varchar(100) NOT NULL,
    is_active bool NOT NULL DEFAULT true,
	created_by varchar(100) NOT NULL,
	created_date timestamptz NOT NULL DEFAULT Now(),
	modified_by varchar(100) NULL,
	modified_date timestamptz NULL DEFAULT Now(),
	deleted_by varchar(100) NULL,
	deleted_date timestamptz NULL DEFAULT '1753-07-01 00:00:00+07:07:12'::timestamp with time zone,
	CONSTRAINT user_role_pk PRIMARY KEY (id),
	CONSTRAINT user_role_un UNIQUE (code),
	CONSTRAINT user_role_un_del UNIQUE (code, deleted_date)
);