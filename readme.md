# Bookings and Making Rservations using Golang

This is a repositiry for go crud with postgres project writing in Golang 

- Go version 1.16

and following external packages :

- Postgres [Postgres driver](https://github.com/lib/pq)
- Routing [Gorilla](https://github.com/gorilla/mux) 
- [joho/godotenv](https://github.com/joho/godotenv) to read environment file  
- [ElephantSQL](https://www.elephantsql.com/) Used to set up the Postgres database

- After starting the app, the CRUD methods could be tested by any API client eg Postman using

- Post crete new user http://localhost:8080/api/newuser with a jason object "name":"string, "age": int, "location": "string"

- Get a user http://localhost:8080/api/user/{id}

- Get all user http://localhost:8080/api/user

- Put update a user http://localhost:8080/api/user/{id}

- Delete a user http://localhost:8080/api/deleteuser/{id}