Bu proqramı istifadə etmək üçün aşağıdakıları etmək lazımdır:

1. İlk öncə Go proqramını download və install edirik. Download üçün link https://golang.org/doc/install
  
2. Daha sonra build edirik: go build ExcelToSql.go ve ya 1 ve 2-ci stepi etmeden ExcelToSql.exe faylini birbasa istifade ede bilersiniz.

3. Daha sonra aşağıdakı komandanı yerinə yetiririk:
$ ExcelToSql.exe -f="fileName" -dn="dictName" -dc="dictCode"
Burada -f file path-i,-dn dictionary-nin adını, -dc dictionary code-nu bildirir.
Example: ExcelToSql.exe -f=message-type.xlsx -dn=ediDict -dc=messageType. 
Bir nece fayl, dictionary name ve dictionary code-da mumkundur. 
Dictionaryler arasinda elaqe varsa qeyd edilme ardicilligi parentden child-a dogru olmalidir.
Example: ExcelToSql.exe -dc dictCode -dn dictName -f parentfile -dc dictCode1 -dn dictname1 -f childFile
    
4. Əgər istərsəniz ExcelToSql.exe faylının path-ini Environment Variables-e əlavə edərək istənilən yerdə istifadə edə bilərsiniz.


