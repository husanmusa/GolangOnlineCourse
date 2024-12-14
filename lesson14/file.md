# database nima?

# DBMS - Database Managment Systems

# RDMBS <- | -> NoSQL

# Relational DBMS

# SQL 

# creata database
 - create database <database_name>

# create table
 - create table <table_name>
 - create table car(id int, name varchar deault '', company varchar(64), year smallint, is_new bool);

 insert into car( name, company, year, is_new) values( 'CyberTruck', 'Tesla', 2023, true);
 insert into car(id, name, company, year, is_new) values(4, 'Matiz', 'Chevrolet', 2005, false);
 insert into car( name, company, year, is_new) values( 'Spark', 'Chevrolet', 2024, true);,(4, 'Chazon', 'BYD', 2022, false);

 - select name, company, year from car where company='chevrolet';

 # update data
  - update car set year=2014, is_new=false where name='Spark';

  # delete table;
  - drop table <table_name>;

# erase data from table;
 - truncate <table_name>;

# change or update table information;
 - alter ...
 - alter table car alter column name TYPE varchar(128);
 - ALTER TABLE car ALTER COLUMN year SET DEFAULT 2024;

 insert into car(id, name, company) values(5, 'CARBAR', 'Nothing');
# 1, 2, 3, 4, 5, 6
 create table car(id serial unique, name varchar default '', company varchar(64) default '', year smallint default 2024, is_new bool default false);
insert into car() values();