-- Sabquery by query ichida bir nechta query lar
-- bo'lsa bu subquery dir. 

-- Subquery lar 4 turga bo'linadi
-- 1. Scalar Subqueries
-- 2. Row Subqueries
-- 3. Column Subqueries
-- 4. Table Subqueries

-- Scalar Sabquery by query ichiqagi subquery dan 
-- faqatgina bitta malumot keladigan subquerydir.
select column_name
    from table1
    where column_name operator
           (select column_name
            from table_name 
            where condition);


-- Row Sabquery by query ichiqagi subquery dan 
-- bir nechta column lar va row lar dagi malumotlar
-- keladigan subquerydir. Keladigan column lar va 
-- row lar dagi malumotlar ko'p bo'lishi yoki kam
-- bo'lishi mumkin va bu Row Sabquery ichidagi where
-- dagi contion ga bog'liq.
 select column_name
    from table1
    where column_name
          in (select column_name
              from table_name 
              where condition);

-- Column Sabquery by query ichiqagi subquery dan 
-- faqatgina column lar dagi malumotlarni keladigan
-- subquerydir, row lardagini emas.
select column_name
    from table1
    where (select column_name
            from table_name 
            where condition);

-- Table Sabquery by query ichiqagi subquery dan 
-- tadval k'rinishida malumotlarni keladigan
-- subquerydir yani bitta table ni nomi yozmasdan
-- u table ni malumotlarini sabquery qilib olib kelish
-- uchun ishlatiladi. Table nomini yozmasdan o'sha table
-- ni olib kelishning bir yo'li.
select column_name
    from
        (select column_name
         from   table1)
    where  condition;


-- yana Correlated Subqueries ham bor. bu subquery
-- tashqi sabquery ning malumotini o'zini ichida foydalana
-- oladigan subquery dir yani tashqi query dagi malumot
-- bilan ichiki query dagi malumotni solishtirish orqali
-- biror bir malumotni olib kelish uchun ishlatish 
-- mumkin.
select e1.employee_name, e1.salary
from employee e1
where salary > 
   (select AVG(salary)
   from employee e2
   where e1.department = e2.department); 