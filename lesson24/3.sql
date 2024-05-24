create database books;

create table author
(
    id  int primary key,
    name varchar not null 
);

create table book
(
    id          int primary key,
    name        varchar not null,
    page        int4,
    author_id   int,
    foreign key (author_id) references author(id)
);

----------------------------------------------------------------------------

insert into author(
    id, name
) values(
    1,'Alisher Navoi'
);

insert into author(
    id, name
) values(
    2,'Abdulla Qodiriy'
);

insert into author(
    id, name
) values(
    3,'Cholpon'
);

insert into author(
    id, name
) values(
    4,'Gafur Ghulom'
);

insert into author(
    id, name
) values(
    5,'Abdulla Qahhor'
);

insert into author(
    id, name
) values(
    6,'Hamid Olimjon'
);

insert into author(
    id, name
) values(
    7,'Erkin Vohidov'
);

insert into author(
    id, name
) values(
    8,'Asqad Mukhtar'
);

-------------------------------------------------------------------

insert into book(
    id, name, page, author_id
) values (
    1, 'Khamsa', 2500, 1
);

insert into book(
    id, name, page, author_id
) values (
    2, 'O''tkan Kunlar', 400, 2
);

insert into book(
    id, name, page, author_id
) values (
    3, 'Mehrobdan Chayon', 300, 2
);

insert into book(
    id, name, page, author_id
) values (
    4, 'Kecha va Kunduz ', 350, 3
);

insert into book(
    id, name, page, author_id
) values (
    5, 'Shum Bola', 250, 4
);

insert into book(
    id, name, page, author_id
) values (
    6, 'Menim O''g''rigina Bolam', 150, 4
);

insert into book(
    id, name, page, author_id
) values (
    7, 'Sarob', 300, 5
);

insert into book(
    id, name, page, author_id
) values (
    8, 'Sinchalak', 200, 5
);

insert into book(
    id, name, page, author_id
) values (
    9, 'Zaynab va Omon', 200, 6
);

insert into book(
    id, name, page, author_id
) values (
    10, 'Oygul bilan Bahtiyor', 280, 6
);

insert into book(
    id, name, page, author_id
) values (
    11, 'Ruhlar isyoni', 250, 7
);

insert into book(
    id, name, page, author_id
) values (
    12, 'Qo''shiqlarim sizga', 150, 7
);

insert into book(
    id, name, page, author_id
) values (
    13, 'Chinor', 300, 8
);

insert into book(
    id, name, page, author_id
) values (
    14, 'Sog''inch', 250, 8
);

insert into book(
    id, name, page, author_id
) values (
    15, 'Turkiston', 200, 8
);

-------------------------------------------------------------------