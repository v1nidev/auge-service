-- object: public.person | type: TABLE --
-- DROP TABLE IF EXISTS public.person CASCADE;
CREATE TABLE public.person (
	id uuid NOT NULL,
	name text NOT NULL,
	CONSTRAINT person_id_pkey PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.person OWNER TO admin;
-- ddl-end --

-- object: public.member | type: TABLE --
-- DROP TABLE IF EXISTS public.member CASCADE;
CREATE TABLE public.member (
	sex character(1) NOT NULL,
	email character varying(255) NOT NULL,
	id_membership_status numeric(2) NOT NULL,
-- 	id uuid NOT NULL,
-- 	name text NOT NULL,
	CONSTRAINT member_sex_check CHECK (sex in ('M', 'F')),
	CONSTRAINT member_pk PRIMARY KEY (id)
)
 INHERITS(public.person);
-- ddl-end --
ALTER TABLE public.member OWNER TO admin;
-- ddl-end --

-- object: public.phone | type: TABLE --
-- DROP TABLE IF EXISTS public.phone CASCADE;
CREATE TABLE public.phone (
	id_person uuid NOT NULL,
	phone_number character varying(20) NOT NULL,
	CONSTRAINT phone_person_id_phone_pk PRIMARY KEY (id_person,phone_number)
);
-- ddl-end --
ALTER TABLE public.phone OWNER TO admin;
-- ddl-end --

-- object: public.emergency_phone | type: TABLE --
-- DROP TABLE IF EXISTS public.emergency_phone CASCADE;
CREATE TABLE public.emergency_phone (
	phone_number character varying(20),
	name text,
	id_member uuid NOT NULL,
	CONSTRAINT emergency_phone_pk PRIMARY KEY (id_member)
);
-- ddl-end --
ALTER TABLE public.emergency_phone OWNER TO admin;
-- ddl-end --

-- object: member_fk | type: CONSTRAINT --
-- ALTER TABLE public.emergency_phone DROP CONSTRAINT IF EXISTS member_fk CASCADE;
ALTER TABLE public.emergency_phone ADD CONSTRAINT member_fk FOREIGN KEY (id_member)
REFERENCES public.member (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: emergency_phone_uq | type: CONSTRAINT --
-- ALTER TABLE public.emergency_phone DROP CONSTRAINT IF EXISTS emergency_phone_uq CASCADE;
ALTER TABLE public.emergency_phone ADD CONSTRAINT emergency_phone_uq UNIQUE (id_member);
-- ddl-end --

-- object: public.membership_status | type: TABLE --
-- DROP TABLE IF EXISTS public.membership_status CASCADE;
CREATE TABLE public.membership_status (
	id numeric(2) NOT NULL,
	status character varying(20) NOT NULL DEFAULT 'inactive',
	CONSTRAINT membership_status_pk PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.membership_status OWNER TO admin;
-- ddl-end --

-- object: membership_status_fk | type: CONSTRAINT --
-- ALTER TABLE public.member DROP CONSTRAINT IF EXISTS membership_status_fk CASCADE;
ALTER TABLE public.member ADD CONSTRAINT membership_status_fk FOREIGN KEY (id_membership_status)
REFERENCES public.membership_status (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: public.employee | type: TABLE --
-- DROP TABLE IF EXISTS public.employee CASCADE;
CREATE TABLE public.employee (
	birthday date,
-- 	id uuid NOT NULL,
-- 	name text NOT NULL,
	CONSTRAINT employee_pk PRIMARY KEY (id)
)
 INHERITS(public.person);
-- ddl-end --
ALTER TABLE public.employee OWNER TO admin;
-- ddl-end --

-- object: public.package_type | type: TABLE --
-- DROP TABLE IF EXISTS public.package_type CASCADE;
CREATE TABLE public.package_type (
	id numeric(3) NOT NULL,
	package_type_description character varying(30) NOT NULL,
	CONSTRAINT package_type_pk PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE public.package_type OWNER TO admin;
-- ddl-end --

-- object: public.package | type: TABLE --
-- DROP TABLE IF EXISTS public.package CASCADE;
CREATE TABLE public.package (
	price numeric(10,2) NOT NULL,
	validity_in_months numeric(3) NOT NULL,
	id_package_type numeric(3) NOT NULL,
	CONSTRAINT package_pk PRIMARY KEY (id_package_type,validity_in_months)
);
-- ddl-end --
ALTER TABLE public.package OWNER TO admin;
-- ddl-end --

-- object: public.member_signs_package | type: TABLE --
-- DROP TABLE IF EXISTS public.member_signs_package CASCADE;
CREATE TABLE public.member_signs_package (
	id_member uuid NOT NULL,
	id_package_type numeric(3) NOT NULL,
	validity_in_months numeric(3) NOT NULL,
	CONSTRAINT member_signs_package_pk PRIMARY KEY (id_member,id_package_type,validity_in_months)
);
-- ddl-end --
ALTER TABLE public.member_signs_package OWNER TO admin;
-- ddl-end --

-- object: phone_person_id_fk | type: CONSTRAINT --
-- ALTER TABLE public.phone DROP CONSTRAINT IF EXISTS phone_person_id_fk CASCADE;
ALTER TABLE public.phone ADD CONSTRAINT phone_person_id_fk FOREIGN KEY (id_person)
REFERENCES public.person (id) MATCH SIMPLE
ON DELETE CASCADE ON UPDATE CASCADE;
-- ddl-end --

-- object: id_package_type_fk | type: CONSTRAINT --
-- ALTER TABLE public.package DROP CONSTRAINT IF EXISTS id_package_type_fk CASCADE;
ALTER TABLE public.package ADD CONSTRAINT id_package_type_fk FOREIGN KEY (id_package_type)
REFERENCES public.package_type (id) MATCH SIMPLE
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: package_fk | type: CONSTRAINT --
-- ALTER TABLE public.member_signs_package DROP CONSTRAINT IF EXISTS package_fk CASCADE;
ALTER TABLE public.member_signs_package ADD CONSTRAINT package_fk FOREIGN KEY (id_package_type,validity_in_months)
REFERENCES public.package (id_package_type,validity_in_months) MATCH FULL
ON DELETE NO ACTION ON UPDATE CASCADE;
-- ddl-end --

-- object: id_member_fk | type: CONSTRAINT --
-- ALTER TABLE public.member_signs_package DROP CONSTRAINT IF EXISTS id_member_fk CASCADE;
ALTER TABLE public.member_signs_package ADD CONSTRAINT id_member_fk FOREIGN KEY (id_member)
REFERENCES public.member (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE CASCADE;
-- ddl-end --


