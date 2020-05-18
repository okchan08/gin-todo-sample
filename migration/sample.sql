drop database TODO_SAMPLE;
create database TODO_SAMPLE CHARACTER SET utf8;

CONNECT TODO_SAMPLE;
create user if not exists 'test'@'%' identified by 'password';
grant all privileges on TODO_SAMPLE.* TO 'test'@'%';

create table if not exists USERS
(
  user_id   serial primary key,
  email     varchar(255) unique not null,
  user_name varchar(255)        not null,
  password  varchar(255)        not null,
  created_at DATETIME   default CURRENT_TIMESTAMP
);

INSERT INTO USERS(EMAIL, USER_NAME, PASSWORD)
VALUES('test1@example.com', 'test1 user', 'test1_pass');

INSERT INTO USERS(EMAIL, USER_NAME, PASSWORD)
VALUES('test2@example.com', 'test2 user', 'test2_pass');
