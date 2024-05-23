create database Users;

\c users;

create table users(ID serial primary key, Name varchar(50), Surname Varchar(50), Age int);

insert into users(Name, Surname, Age) values('Saidakbar', 'Pardaboyev', 19);
insert into users(Name, Surname, Age) values('Muhammadjon', 'Ko''palov', 19),('Iskandar', 'Iskandarov', 22);
insert into users (Name, Surname, Age) values('Akrom', 'Iskandarov', 19),('Musohon', 'Mo''minov', 28);
insert into users (Name, Surname, Age) values('Umar', 'Hattob', 45),('Xadicha', 'Xuvaylid', 30),('Oysha', 'Abu Bakr', 25);
insert into users (Name, Surname, Age) values('Akbar', 'Mamarasulov', 21),('Zokir', 'Isroilov', 12);

