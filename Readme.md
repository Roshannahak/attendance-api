# ANTRY-API
Project **Antry** for *Visitor*, *employee*, *student* entry log managament system.

## Pull Docker Image
first you have to ensure that install docker on your system after that go to the terminal and write down code. API server port `:5252`
```
docker pull roshannahak/antry_api_server
```
```
docker run -p 5252:5252 roshannahak/antry_api_server
```

## Server build file
Linux | [download](https://github.com/Roshannahak/attendance-api/blob/main/main)
--- | ---
**Windows** | [download](https://github.com/Roshannahak/attendance-api/blob/main/main.exe)

## API Documentation :-

### BASE URL : https://(current ip):5252/api

Method | Routes    | Description | Request Body
------ | --------- | ----------  | -----------
GET    | /stats    | get dashboard stats |
POST   | /auth/admin/register | admin registration |
POST   | /auth/admin/login | admin login |
GET   | /auth/admin/decode | decrypt admin token | Header({"x-auth-admin": token})
POST   | /auth/student/register | student registration |
POST   | /auth/student/login | student login |
GET   | /auth/student/decode | decrypt student token | Header({"x-auth-student": token})
POST   | /auth/visitor | visitor login |
GET   | /visitor | get all visitor list |
DELETE   | /visitor/:visitorId | delete visitor by id |
GET   | /visitor/:quary | search visitor by name |
GET   | /visitor/log | get all visitor logs list |
POST   | /visitor/log/entry | visitor check-in check-out |
GET | /visitor/log/checkin | get check-in visitor list |
GET | /visitor/log/id/:visitorId | get particular visitor all log list by visitor id |
GET | /visitor/log/:logId | get visitor log by logId |
GET | /student | get all student list |
DELETE | /student/:studentId | delete student by id |
GET | /student/:quary | search student by name |
GET | /student/log | get all student log list |
POST | /student/log/entry | student check-in check-out |
GET | /student/log/checkin | get student check-in list |
GET | /student/log/id/:studentId | get particular student all log list by student id |
GET | /student/log/:logId | get student log by logId |
GET | /admin | get all admin list |
GET | /room | get all rooms |
POST | /room | add rooms |
DELETE | /room/:roomId | delete room by room id |
PUT | /room/:roomId | update room by room id |



