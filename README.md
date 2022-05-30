# How to run
In the terminal, run below statement:
<br> docker compose up

# Code structure
## Used packages:
### api
Build api services to create/query data from postgres.
<br> - I used sqlc library to implement create, read operations, it generate golang script from SQL script.
<br> - Gin to implement HTTP API
### db
It stores procedures to run database migration and generate go script from script generated by sqlc.

### util
Load config environment variables from app.env to api service.
random.go simply generates random data for testing purposes.

### Test api with postmain
GET: http://localhost:8080/multipleorder #Auto-generate sample data & get 5 sample records.
<br> POST: http://localhost:8080/order #Create singular record {"user":string, "product":string, "amount":int}
<br> GET: http://localhost:8080/order

