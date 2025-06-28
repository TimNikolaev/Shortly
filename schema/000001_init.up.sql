CREATE TABLE users
(
  id serial not null unique,
  created_at timestamp default null,
  updated_at timestamp default null,
  deleted_at timestamp default null,
  name varchar(255) not null,
  email varchar(255) not null unique,
  password_hash varchar(255) not null
);

CREATE TABLE links
(
  id serial not null unique,
  created_at timestamp default null,
  updated_at timestamp default null,
  deleted_at timestamp default null,
  user_id int references users (id) on update cascade on delete set null,
  url varchar(255) not null,
  hash varchar(255) not null
);

CREATE TABLE stats
(
  id serial not null unique,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp,
  link_id int references links (id) on update cascade on delete set null,
  clicks int not null default 0,
  date_stat date not null
); 