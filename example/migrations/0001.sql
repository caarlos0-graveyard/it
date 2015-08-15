CREATE SEQUENCE books_seq;
CREATE TABLE books (
  id bigint not null primary key,
  name character varying(200) not null
);
