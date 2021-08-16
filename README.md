Versi Golang = go1.14.5
FrameWork = Echo
Platform API = Postman
postman collection = UserManagement_collection.json
Database = MySQL
File SQL Database = user_manag.sql (file sql)
Database name = user_manag
Database username = root
Database password = ""


How to use:
-----------
Buat Database di MySQL dengan nama user_manag, dan import file SQL (user_manag) pada database 
Import postman collection (UserManagement_collection.json), sudah terdapat body dan parameter untuk pengisian

Run Application
---------------
Run Terminal pada path direktori : go run main.go


1. Input User
postman: 
- Create User : ("/user)
req body:
{
"username" : "username_afina",
"password" : "password22",
"name" : "name_afina"
}

Note :
- Aplikasi sudah memiliki validasi ketika input username < 3 karakter, password < 7 karakter, dan username <3 karakter, dan password yang tersimpan sudah di hash menggunakan becrypt


3. View User
postman: 
- Read All with pagination : ("/user/:limit/:offset")
- Read By Id User : ("/user/:userId")
  

4. Update User
postman: 
- Update User : ("/user/:userId)
req body:
{
"username" : "username_afina12",
"password" : "password22",
"name" : "name_afina12"
}

Note :
- Aplikasi sudah memiliki validasi ketika input username < 3 karakter, password < 7 karakter, dan username <3 karakter, dan password yang tersimpan sudah di hash menggunakan becrypt

5. Delete User
postman: 
- Delete User : ("/user/:userId)
