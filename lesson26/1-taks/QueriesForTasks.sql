-- Printing Student with their average grade in their each courses
select
    s.name as student,
    c.name as course,
    round(avg(grade)::numeric, 1) as average_grade
from
    courses as c 
left join
    students_courses as sc 
        on c.id = sc.student_course_id
left join
    students as s 
        on s.id = sc.student_id
inner join
    grade as g 
        on g.student_course_id = sc.id
group by
    c.name, s.name
;

-- Task 2 -------------------------------------------------------------
with averageGradeOfEachStudents as (
    select
        s.id as student_id,
        c.id as course_id,
        round(avg(grade)::numeric, 1) as average_grade
    from
        courses as c 
    left join
        students_courses as sc 
            on c.id = sc.student_course_id
    left join
        students as s 
            on s.id = sc.student_id
    inner join
        grade as g 
            on g.student_course_id = sc.id
    group by
        c.id, s.id
),

maxAverageGradeOfPerCourse as (
    select
        course_id,
        max(average_grade) as max_average
    from
        averageGradeOfEachStudents
    group by
        course_id
)

select
    c.name as course,
    array_agg(s.name) as students,
    a.average_grade as max_average_grade
from
    students as s
left join
    students_courses as sc
        on s.id = sc.student_id
left join
    courses as c 
        on sc.student_course_id = c.id
inner join
    averageGradeOfEachStudents as a 
        on c.id = a.course_id and s.id = a.student_id
inner join
    maxAverageGradeOfPerCourse as m
        on c.id = m.course_id and m.max_average = a.average_grade
group by
    c.id, c.name, a.average_grade
;

---------------------------------------------------------------------

-- Printing Students with their courses
select s.name as student_name, c.name as course
from students as s
left join students_courses as sc
on s.id = sc.student_id
left join courses as c
on sc.student_course_id = c.id;

-- Printing course's students and their grades
select
    c.name,
    array_agg(distinct s.name) as students,
    array_agg(
        case
            when g.grade is not null then g.grade end
    ) as average_grade
from courses as c 
left join students_courses as sc 
on c.id = sc.student_course_id
left join students as s
on sc.student_id = s.id
left join grade as g 
on g.student_course_id = sc.id
group by c.name;

-- Printing Average grade in each courses
select
    c.name,
    sum(grade) / count(grade) as average_grade
from courses as c 
left join students_courses as sc 
on c.id = sc.student_course_id
left join students as s
on sc.student_id = s.id
left join grade as g 
on g.student_course_id = sc.id
group by c.name;

---------------------------------------------------------------

-- Printing students with their ages in each courses
select
    c.name as course,
    array_agg((s.name, s.age)) as student
from
    courses as c
left join
    students_courses as sc on sc.student_course_id = c.id 
left join
    students as s on sc.student_id = s.id
group by
    c.name
;

-- Printing Yongest students in each courses
with yongest_age_in_each_course as (
    select 
        min(age) as min_age,
        c.id as course_id
    from 
        courses as c 
    left join
        students_courses as sc on sc.student_course_id = c.id 
    left join
        students as s on s.id = sc.student_id
    group by
        c.id
)

select
    c.name as course,
    array_agg((s.name, s.age)) as student
from
    courses as c
left join
    students_courses as sc on sc.student_course_id = c.id 
left join
    students as s on sc.student_id = s.id
inner join
    yongest_age_in_each_course as y 
        on y.min_age = s.age and y.course_id = c.id
group by
    c.name
order by
    c.name
;

---------------------------------------------------------------

-- Printing the best group in studying
with average_grades_list as (
    select
        sum(grade) / count(grade) as average_grade
    from 
        courses as c 
    left join 
        students_courses as sc 
            on c.id = sc.student_course_id
    left join 
        students as s
            on sc.student_id = s.id
    left join 
        grade as g 
            on g.student_course_id = sc.id
    group by
        c.id
),

max_average_grade as (
    select
        max(average_grade) as MaxAvg
    from 
        average_grades_list
)

select
    c.name as course_name,
    sum(grade) / count(grade) as average_grade
from 
    courses as c 
left join 
    students_courses as sc 
        on c.id = sc.student_course_id
left join 
    students as s
        on sc.student_id = s.id
left join 
    grade as g 
        on g.student_course_id = sc.id
group by
    c.name
having
    sum(grade) / count(grade) = (select MaxAvg from max_average_grade)
;    