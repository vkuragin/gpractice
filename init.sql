-- create db and user
create database if not exists gpractice;
create user if not exists gpractice identified by '123';
grant all privileges on gpractice.* to gpractice@'%';

-- create tables
create table gpractice.practice (
  id bigint not null auto_increment primary key,
  date date not null,
  duration bigint not null default 0
) engine=InnoDB;