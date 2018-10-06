create table users (
  id integer primary key,
  name text unique,
  password text
)

create table reactions (
  from integer,
  to integer,
  type integer
)
