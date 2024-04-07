create database test_crud

create table seller (
    "id" serial primary key,
    "name" character varying(250),
    "suname" character varying(250)
);
insert into seller(name, suname) values('aman', 'amanow'),('mered', 'mredow')

create table "user" (
    "id" serial primary key,
    "name" character varying(250),
    "surname" character varying(250),
    "password" character varying(250)
);

CREATE TABLE "workers" (
    "id" serial primary key,
    "user_id" integer not null,
    "role" character varying(100),
    "page_id" integer
);
