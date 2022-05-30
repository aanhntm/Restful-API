# How to run
In the terminal, run below statement:
<br> docker compose up

# Code structure
## Used packages:
### api
Build api services to create/query data from postgres.
<br> - I used sqlc library to implement create, read operations, it generate golang script from SQL script.
<br> - Gin to implement HTTP API.
### db
It stores procedures to run database migration and generate go script from script generated by sqlc.

### util
Load config environment variables from app.env to api service.
random.go simply generates random data for testing purposes.

# Demo API
### Testing api with postman & results:

The service is up and running:
![Screen Shot 2022-05-30 at 13 32 16](https://user-images.githubusercontent.com/106065029/170930852-760cc7b0-56fa-4d55-842a-049d1a3d59bd.png)

<br> GET: http://localhost:8080/multipleorder #Auto-generate sample data & get 5 sample records.

![Screen Shot 2022-05-30 at 13 26 06](https://user-images.githubusercontent.com/106065029/170930060-ec1a8a00-b57b-4848-ae92-220241ade09b.png)

<br> POST: http://localhost:8080/order #Create singular record {"user":string, "product":string, "amount":int}
![Screen Shot 2022-05-30 at 13 28 08](https://user-images.githubusercontent.com/106065029/170930332-a02f2be4-ff3c-481b-9cc0-c8ec5f3df5bc.png)

<br> GET: http://localhost:8080/order

![Screen Shot 2022-05-30 at 13 20 04](https://user-images.githubusercontent.com/106065029/170930210-94c3edd6-fc01-4c13-a59a-6644540f04ec.png)

<br> Checking Postgres database:
<img width="1430" alt="Screen Shot 2022-05-30 at 13 29 07" src="https://user-images.githubusercontent.com/106065029/170930470-f3859ac7-1ff7-4d19-b024-d9ea9b824a13.png">

<br> Checking result generated by terminal:
![Screen Shot 2022-05-30 at 13 30 12](https://user-images.githubusercontent.com/106065029/170930599-64fb1352-b9d8-469f-ba1e-fbb97eadcba2.png)
