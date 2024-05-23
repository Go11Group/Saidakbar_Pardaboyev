-- update bu jadvalga kiritilgan biror bir malumotni o'zgartirish uchun kerak.
-- Misol uchun biror bir ism xato kiritilgan bo'lsa uni o'zgartisish uchun update dan foydalanasiz.
-- buning uchun o'sha qatorda bor lekin boshqa qatorlarda bo'lmagan biror bir malumot orqali
-- o'zgartirasiz.
update users set Name = 'Saidakbar' where ID = 10;

-- Agar bir biror bir malumotni table dan o'chirmoqchi bo'lsak delete dan foydalamiz.
-- buning uchun ham update kabi o'chirmoqchi bo'lgan malumotimiz turgan qatordagi unique 
-- malumot orqali o'chiramiz.
delete from users where ID = 10;

-- Group by ni biz qachonki biror bir columnda turgan malumotlar orqali guruhlab olib
-- biror bir ish qilmoqchi bo'lsak ishlatamiz, misol uchun har bir bo'limda nechda
-- ishchi borligini bilish uchun ishlatamiz.
-- Group by ni biz odatda COUNT(), MAX(), MIN(), SUM(), AVG() kabi lar bilan ishlatadiz.
select COUNT(*) from uzum group by department;

-- data base dan malumotlarni malum bir tartib da olish uchun ORDER BY dan foydalanamiz.
-- Qaysi ustun orqali sortlab chiqarilishini ham so'rab olamiz.
select * from users ORDER BY age;

-- Tasavvur qilish agar bizda bir javdal bor. Unda juda ham ko'p ustunlar bor. Ustunlari
-- juda ko'p bo'lgan jadvalni tushunish qiyin bo'ladi shuning uchun u jadvalni jadvalchalarga
-- bo'lib tashlanadi. Agar bizga bo'lingan jadvallarni 2 tasini vaqtincha birlashtirish kerak 
-- bo'lib qolsa biror bir vazifani bajarish uchun Join dan joydalanasiz. Join 4 turga bo'linadi.
-- INNER JOIN, LEFT JOIN, RIGHT JOIN, FULL JOIN larga. Join ni ishlatish uchun biror bir ustun
-- asosida ishlatasiz
-- INNER JOIN bu 2 jadvaldagi bir biri bilan bo'g'liq bo'lgan usuni borlarni birlashtiradi
-- qolganlarini tashlab ketadi.
select testproduct_id, product_name, category_name
from testproducts as t
inner join categories as c on t.category_id = c.category_id;

-- LEFT JOIN bu chap tarafda turgan table ga o'ng tarafdagi table dan bir-biriga 
-- bog'liq bo'lganlarni qo'shadi, chap tarafdagi jadval ning o'ng tarafdagi jadval ga
-- bog'langani bo'lmasa ularga nil tenglashtiradi. LEFT JOIN chap tarafdagi jadvalga
-- o'ng tarafdagi jadvalni qo'shadi agar bir biriga bog'liq joyi bo'lsa.
select testproduct_id, product_name, category_name
from testproducts as t
left join categories as c on t.category_id = c.category_id;

-- RIGHT JOIN bu o'ng tarafda turgan table ga chap tarafdagi table dan bir-biriga 
-- bog'liq bo'lganlarni qo'shadi, o'ng tarafdagi jadval ning chap tarafdagi jadval ga
-- bog'langani bo'lmasa ularga nil tenglashtiradi. RIGHT JOIN o'ng tarafdagi jadvalga
-- chap tarafdagi jadvalni qo'shadi agar bir biriga bog'liq joyi bo'lsa.
select testproduct_id, product_name, category_name
from testproducts as t
right join categories as c on t.category_id = c.category_id;

-- full join 2 table hamma malumotini ni birlashtiradi, bir biri bilan bog'liqligi bo'lmasa
-- ham yani left dagi table ni malumoti right da bo'lmasa nill bilan birlashadi agar teskarisi
-- bo'lsa birlashtiradi. rightni ham huddi shunday qiladi yani left join bilan right join ni 
-- birlashtirilgan variyanti bu full join.
select testproduct_id, product_name, category_name
from testproducts as t
full join categories as c on t.category_id = c.category_id;