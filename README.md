# wallet-ap-graphql
this is a backend of wallet app for personal spending and income management.

technologies used:
golang
graphql
postgres
redis
jwt authorization
smtp
--------------------------------------------------------------------------------------------------------------------------------------------------------------------

appni ishga tushirishdan oldin _setup folderdagi model.sql va triggers.sql fayllarini database'da execute qilib olish kerak. trigger funksiyalar dastur to'gri ishlashi uchun kerak.

query/mutation yozishda muammo bo'lsa asosiy folderda graphql-queries-and-mutations.txt faylida shu dastur uchun yozilgan sample querylar bor.

redis database ishlashi uchun kompyuterda o'rnatilgan bo'lishi kerak. agar o'rnatilmagan bo'lsa, uning o'rniga redis folderidagi redis.go faylida kommentariyaga olib qo'yilgan kodni ishlatish mumkin.

----------------------------------------------------------------------------------------------------------------------------------------------------------------

Before launching the app, you need to execute the "model.sql" and "triggers.sql" files in the _setup folder . The trigger functions are necessary for the program to work properly.

If there is a problem writing query / mutation, the main folder contains sample queries written for this program in the "graphql-queries-and-mutations.txt" file.

redis database must be installed on the computer for operation. if not installed, you can use the code that commented in the "redis.go" file in the redis folder instead.
