insert into students(
    name, surname, age
) values (
    'Saidakbar', 'Pardaboyev', 19
), (
    'Muhammadjon', 'Ko''palov', 19
), (
    'Iskandar', 'Iskandarov', 22
), (
    'Javlon', 'Shamsiyev', 20
), (
    'Ibragim', 'Ibragimov', 25
), (
    'Sharapat', 'Nappiyeva', 35
), (
    'Nappi', 'Sharapateva', 29
);

insert into courses (
    name, description
) values (
    'Golang11', 'Learning Golang for earnging money'
), (
    '.Net', 'Learning .Net for earnging money'
), (
    'Flatter', 'Learning Flatter for earnging money'
), (
    'PHP', 'Learning PHP for earnging money'
);

insert into students_courses (
    student_id, student_course_id
) values (
    'a8f5e97c-e4e0-4763-9626-920d3087549d', '3d5a7fee-e64e-41e8-b01b-4f220f26fdd1'
), (
    '34777ba5-1936-43f0-80a1-2d8afc84f66a', '3d5a7fee-e64e-41e8-b01b-4f220f26fdd1'
), (
    '07508da0-d15a-4759-87ab-7b09d4f590ed', '3d5a7fee-e64e-41e8-b01b-4f220f26fdd1'
), (
    '705aab3c-5623-4208-a63e-b3a0e6b41552', 'fd5c7166-3ac5-43aa-b8a9-1abc0afef6f1'
), (
    '58dc8905-e9d8-463a-a858-8d704a0bb212', 'fd5c7166-3ac5-43aa-b8a9-1abc0afef6f1'
), (
    'd4247727-1cc2-4c25-8221-d66b0687f425', 'a0715723-6558-47ef-8980-95e8037a04c1'
), (
    '3b1b20d8-45e5-4707-b195-753222c3f431', 'a0715723-6558-47ef-8980-95e8037a04c1'
);

insert into grade
    (grade, student_course_id)
values
    (5, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (4, '21c7e4cd-42a4-4b06-9ce0-3134fd88d14a'),
    (3, 'f511837d-4c8e-44e3-87dd-c8118f22cade'),
    (5, 'a6d673f5-affb-4070-995a-945fabe90929'),
    (4, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (2, '21c7e4cd-42a4-4b06-9ce0-3134fd88d14a'),
    (1, 'f511837d-4c8e-44e3-87dd-c8118f22cade'),
    (5, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (4, 'a6d673f5-affb-4070-995a-945fabe90929'),
    (5, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (4, '21c7e4cd-42a4-4b06-9ce0-3134fd88d14a'),
    (3, 'f511837d-4c8e-44e3-87dd-c8118f22cade'),
    (3, '6d6f28af-e04b-4564-931d-8b83a65c7059'),
    (4, 'a6d673f5-affb-4070-995a-945fabe90929'),
    (4, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (2, '21c7e4cd-42a4-4b06-9ce0-3134fd88d14a'),
    (1, '6d6f28af-e04b-4564-931d-8b83a65c7059'),
    (5, 'f511837d-4c8e-44e3-87dd-c8118f22cade'),
    (5, 'a6d673f5-affb-4070-995a-945fabe90929'),
    (1, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (2, '21c7e4cd-42a4-4b06-9ce0-3134fd88d14a'),
    (3, '6d6f28af-e04b-4564-931d-8b83a65c7059'),
    (4, 'a6d673f5-affb-4070-995a-945fabe90929'),
    (4, '6d6f28af-e04b-4564-931d-8b83a65c7059'),
    (3, 'f511837d-4c8e-44e3-87dd-c8118f22cade'),
    (2, 'a6d673f5-affb-4070-995a-945fabe90929'),
    (1, '21c7e4cd-42a4-4b06-9ce0-3134fd88d14a'),
    (5, '6d6f28af-e04b-4564-931d-8b83a65c7059'),
    (5, 'fcc0d9ba-dbf6-4ad8-b509-0558fb314ca5'),
    (5, 'cfcb48b4-9de3-4760-bfc1-a1ad1e676b41'),
    (5, 'cfcb48b4-9de3-4760-bfc1-a1ad1e676b41'),
    (5, '6d6f28af-e04b-4564-931d-8b83a65c7059'),
    (5, 'f511837d-4c8e-44e3-87dd-c8118f22cade');
    (3, '6d6f28af-e04b-4564-931d-8b83a65c7059');

