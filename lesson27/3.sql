-- .Query bu database dagi biror bir table dagi malumotlarni olib keish uchun 
-- ishlatiladi yani malumotlarni jadval ko'rinishida olib kela olamiz.
-- .QueryRow dan farqi .QueryRow dan farqi .QueryRow table dan faqat bir qator
-- dagi malumotlarni olib kela oladi. .Query esa table dagi barcha qatorlardagi
-- malumotlarni olib kela oladi.

-- yozilishi structurasi
rows, err := db.Query("select name, age from students")
if err != nil {
    panic(err)
}
defer rows.Close()
var name string
var age int

for rows.Next() {
    err := rows.Scan(&name, &age)
    if err != nil (
        panic(err)
    )
    fmt.Println(name, age)
}

err = rows.Err()
if err != nil {
    panic(err)
}

-- Table dagi malumotlar *sql.Rows type li rows o'zgaruvchiga
-- malumotlar keladi. Birinchi error malumotlarni rows nomli
-- o'zgaruvchiga yozyotganida xatolik chiqsa tekshirish uchun.
-- rows nomi o'zgaruvchi oxirida yopilishi kerak. Malumotlarni
-- rows dan qatorma-qator o'qisa bo'ladi. Buning uchun .Next()
-- funksiyasi orqali rows dan o'qilmagan malumot qolganmi yoki
-- yo'qligini tekshiramiz for da. va for ni ichida .Scan() 
-- funksiyasi orqali bitta-bitta malumotni o'qib olamiz.
-- .Scan() ham error qaytaradi va shuni ham tekshirib oladiz.
-- oxirgi .Err() dan qaytadigan error bilan .Scan() funksiyasidan
-- qaytadigan errorlarni farqi .Scan() kelyotgan malumotni 
-- berilgan o'zgaruvchilarga joylashtirishda chiqgan xatolarni
-- qaytaradi misol uchun type i xato bo'lsa. .Err() esa
-- rows dan kelyotgan malumotda xatolik bo'lsa error qaytaradi.

