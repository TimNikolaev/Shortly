CREATE TABLE users
(
  id serial not null unique,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
);

CREATE TABLE links
(
  id serial not null unique,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  user_id int references users (id) ON UPDATE CASCADE ON DELETE SET NULL,
  url varchar(255) not null,
  hash varchar(255) not null
);

CREATE TABLE stats
(
  id serial not null unique,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  link_id int references links (id) ON UPDATE CASCADE ON DELETE SET NULL,
  clicks int not null default 0,
  date_stat date not null
); 