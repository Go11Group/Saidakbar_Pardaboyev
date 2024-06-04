create database employees;

create table employees(
    id uuid default gen_random_uuid() not null,
    firstname varchar not null,
    surname varchar not null,
    age int not null,
    salary int not null,
    experience int not null
)

create index employees_index on employees(id);
create index employees_index on employees using hash (id);

create index employees_index on employees(firstname);
create index employees_increate index employees_index on employees(salary, firstname, id);
dex on employees using hash (firstname);

create index employees_index on employees(salary);
create index employees_index on employees using hash (salary);

create index employees_index on employees(id, firstname, salary);
create index employees_index on employees(firstname, id, salary);
create index employees_index on employees(salary, firstname, id);

explain (analyse)
select * from employees
where id = '1c236221-7c8f-4e5d-b1a3-1b822de281ab';

explain (analyse)
select * from employees
where salary = 999999;

explain (analyse)
select * from employees
where firstname = 'Lela';

explain (analyse)
select * from employees
where id = '1c236221-7c8f-4e5d-b1a3-1b822de281ab' and 
firstname = 'Jedediah' and salary = 999999;

explain (analyse)
select * from employees
where firstname = 'Jedediah' and id = '1c236221-7c8f-4e5d-b1a3-1b822de281ab'
and salary = 999999;

explain (analyse)
select * from employees
where salary = 999999 and id = '1c236221-7c8f-4e5d-b1a3-1b822de281ab'
and firstname = 'Jedediah';

drop index employees_index;