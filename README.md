# short-url

##Folders description

**Folder with project specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_, _utils_ or else.

- `./configs` folder for configuration functions
- `./middleware` folder for add middleware (Fiber and yours)
-  `./controllers` folder with all endpoints controllers
-  `./db` folder with redis db connector
- `./routes` folder for describe routes of your project
- `./repository` folder for describe `const` of your project
- `./helpers` folder with utility functions (jwt, server starter, error checker, etc)

##Request examples

- `GET auth token`

`curl -X 'GET' \
'http://127.0.0.1:3000/api/v1/token/new' \
-H 'accept: application/json'`

- `POST long URL example`

`curl -X 'POST' \
'http://127.0.0.1:3000/api/v1/shorten' \
-H 'accept: application/json' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjcyMTE1NDd9.7USryPsmdME4XjvvjmkZIITz_lEz279Ix6Qs51f0zH8' \
-H 'Content-Type: application/json' \
-d '{
"expiry": 24,
"short": "",
"url": "www.qwwe344.er"
}`

- `GET long URL redirection example`

`curl -X 'GET' \
'http://127.0.0.1:3000/api/v1/resolve/F27kDNkiHmj' \
-H 'accept: application/json'`
