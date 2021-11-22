CREATE TABLE IF NOT EXISTS public.statuses(
  id serial PRIMARY KEY,
  code varchar(100) NOT NULL,
  display_name varchar(100) NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz,
  deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS public.roles(
  id serial PRIMARY KEY,
  code varchar(100) NOT NULL,
  display_name varchar(100) NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz,
  deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS public.role_types(
  id serial PRIMARY KEY,
  code varchar(100) NOT NULL,
  display_name varchar(100) NOT NULL,
  created_at timestamptz NOT NULL,
  updated_at timestamptz,
  deleted_at timestamptz
);


CREATE TYPE provider AS ENUM('OWN', 'COGNITO');

CREATE TABLE IF NOT EXISTS public.users(
  id serial PRIMARY KEY,
  user_slug uuid DEFAULT uuid_generate_v4() UNIQUE NOT NULL,
  email text NOT NULL,
  password text NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  age text,
  birth_date text,
  address text,
  phone_number text,
  access_token text,
  provider provider DEFAULT 'OWN' NOT NULL,
  status_id integer REFERENCES statuses(id),
  role_id integer REFERENCES roles(id),
  role_type_id integer REFERENCES role_types(id),
  last_login timestamptz,
  created_at timestamptz NOT NULL,
  created_by text NOT NULL,
  updated_at timestamptz NOT NULL,
  updated_by text NOT NULL,
  deleted_at timestamptz,
  deleted_by text,

  UNIQUE (email),
  UNIQUE (phone_number),

  CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.statuses (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT role_fk FOREIGN KEY (role_id) REFERENCES public.roles (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT role_types_fk FOREIGN KEY (role_type_id) REFERENCES public.role_types (id) ON DELETE CASCADE ON UPDATE CASCADE
);

INSERT INTO public.statuses(code, display_name, created_at, updated_at) VALUES('PENDING', 'Pending', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.statuses(code, display_name, created_at, updated_at) VALUES('APPROVED', 'Approved', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.statuses(code, display_name, created_at, updated_at) VALUES('BANNED', 'Banned', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;

INSERT INTO public.roles(code, display_name, created_at, updated_at) VALUES('ADMIN', 'Admin', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.roles(code, display_name, created_at, updated_at) VALUES('STAFF', 'Staff', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.roles(code, display_name, created_at, updated_at) VALUES('REPORTER', 'Reporter', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.roles(code, display_name, created_at, updated_at) VALUES('CUSTOMER', 'Customer', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.roles(code, display_name, created_at, updated_at) VALUES('GUEST', 'Guest', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;

INSERT INTO public.role_types(code, display_name, created_at, updated_at) VALUES('OPERATOR', 'Operator', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.role_types(code, display_name, created_at, updated_at) VALUES('USER', 'User', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;
INSERT INTO public.role_types(code, display_name, created_at, updated_at) VALUES('SYSTEM', 'System', CURRENT_TIMESTAMP
, CURRENT_TIMESTAMP) RETURNING id, code, display_name, created_at, updated_at, deleted_at;