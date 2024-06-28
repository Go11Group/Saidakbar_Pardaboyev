create table transports (
    bus_number int primary key not null,
    stations varchar[] not null
);

create table weathers (
    city varchar not null,
    date date not null,
    temperature float not null,
    humidity int null null,
    wind float not null
);