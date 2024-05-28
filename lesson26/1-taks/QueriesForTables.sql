Create database Students_And_Courses;

create table students (
    id uuid primary key not null default gen_random_uuid(),
    name varchar not null,
    surname varchar not null,
    age int not null check(age >= 18)
);

create table courses (
    id uuid primary key not null default gen_random_uuid(),
    name varchar not null,
    description text default ''
);

create table students_courses (
    id uuid primary key not null default gen_random_uuid(),
    student_id uuid not null references students(id),
    student_course_id uuid not null references courses(id)
);

create table grade (
    id uuid primary key not null default gen_random_uuid(),
    student_course_id uuid not null references students_courses(id),
    grade int check(grade >= 1 and grade <= 5)
);