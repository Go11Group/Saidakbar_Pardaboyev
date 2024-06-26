create table book(
    book_id uuid primary key default gen_random_uuid() not null,
    title varchar not null,
    author varchar not null,
    year_published int not null
);

create table borrowing(
    id uuid primary key default gen_random_uuid() not null,
    book_id uuid references book(book_id),
    user_id uuid not null
);