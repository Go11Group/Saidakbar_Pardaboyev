Create database cars

 -------------------------------------------------------

create 
    table users(
        id uuid primary key not null default gen_random_uuid(),
        name varchar not null
    );

create
    table cars(
        id uuid primary key not null default gen_random_uuid(),
        name varchar,
        year int,
        brand varchar
    );

create 
    table connectionOfTwoTable(
        id uuid primary key not null default gen_random_uuid(),
        user_id uuid references users(id),
        car_id uuid references cars(id)
    );

---------------------------------------------------------------------

insert into users(name)
values ('Saidakbar'),('Muhammadjon'),('Iskandar'),('Javohir'),
('Abbos'),('Musoxon'),('Sharapat'),('Nappi'),('Bo''rivoy'),
('Qo''zivoy');

-----------------------------------------------------------

insert into cars(
    name, year, brand
) values (
    'Maluba', 2015, 'Chevrolet'
),(
    'Kaptiva', 2012, 'Chevrolet'
),(
    'Gentra', 2005, 'Chevrolet'
),(
    'Damas', 2000, 'Chevrolet'
),(
    'Spark', 1999, 'Chevrolet'
),(
    'Tiko', 1800, 'Chevrolet'
);

-------------------------------------------------------------

insert into connectionOfTwoTable(
    user_id, car_id
) values (
    '5c1d8802-46b9-4230-92ad-f61da498dba0', '7ce17404-96f4-4041-9612-9c3d0e7bb91e'
),(
    '5c1d8802-46b9-4230-92ad-f61da498dba0', '7484eb94-03cc-47f1-a156-5944badb6db9'
),(
    'df4bbfe5-bfa2-4e39-9d8b-cc80c4c176d7', '06462557-825b-4bd3-a910-cbd2b5cd00a4'
),(
    'df4bbfe5-bfa2-4e39-9d8b-cc80c4c176d7', 'f6638571-68fc-4af2-8897-6c7ff6967b56'
),(
    '4238312e-e4a1-4d6e-a77d-a1adc23aafaf', '06462557-825b-4bd3-a910-cbd2b5cd00a4'
),(
    '4238312e-e4a1-4d6e-a77d-a1adc23aafaf', 'b2bce9d2-2748-4005-9ca3-739d0155dd3f'
),(
    '4238312e-e4a1-4d6e-a77d-a1adc23aafaf', 'ce42f3c6-dfbd-4a3d-8d4d-011f62458b54'
),(
    '3172767b-9b73-49ed-8875-e757a628ee37', '7ce17404-96f4-4041-9612-9c3d0e7bb91e'
),(
    '54463345-0fdf-4d0e-a6ab-49d3a8a48075', '06462557-825b-4bd3-a910-cbd2b5cd00a4'
);

------------------------------------------------------------------------------

create type car_info as (
    name varchar,
    year int,
    brand varchar
);

select u.name, array_agg((c.name, c.year, c.brand)::car_info) as cars
from users as u
right join connectionOfTwoTable as con 
on u.id = con.user_id
left join cars as c
on c.id = con.car_id
group by u.id
order by u.name;

-----------------------------------------------------------------------

select c.name, array_agg(u.name)
from cars as c
right join connectionOfTwoTable as con 
on c.id = con.car_id
left join users as u 
on con.user_id = u.id
group by c.id, c.name
order by c.id;