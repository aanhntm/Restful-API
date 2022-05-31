# How to run
Open the terminal in directory of Restful-API, run below statement:
```
make assignment
```

# Code structure
## Used packages:
### api
Build api services to create/query data from postgres.
<br> - I used ```sqlc``` library to implement Create, Read operations, it generates golang scripts from SQL scripts.
<br> - ```Gin``` to implement HTTP API.
### db
It stores procedures to run database migrations and generates go scripts from scripts generated by ```sqlc```.

### util
Load config environment variables from app.env to api service.
<br> ```random.go``` simply generates random data for testing purposes.

## Dockerfile
Building docker images has been separated to 2 stages due to the large size of images was generated when images were build with only 1 stage.
#### Build stage:
App was built based on image ```golang:1.18.2-alpine3.16``` with the ```/app``` working directory. Then all the files in the current directory will be copied to the image we are going to build. I used command ```RUN go build -o main main.go``` to build a binary executable file. ```RUN curl -L https://github.com/golang-migrate ... ``` was used to build data schema migration and will be run after we install ```curl```.
#### Run stage:
Copy and run all neccessary file from Build stage to build images. ```EXPOSE``` is set to 8080, which means the container will be exposed to the port 8080 at its runtime. **CMD** defines the default command to run when the container starts, it will run the executable file in the Build stage, which is ```/app/main```. ```RUN chmod +x wait-for.sh``` and ```RUN chmod +x start.sh``` is to authorize the permission to run ```wait-for.sh``` (a libraby allow us to build containers in order) and ```start.sh``` (to instruct the docker-compose to build schema migrations before building containers). ```ENTRYPOINT``` helps to run all the steps in order.

## docker-compose.yaml
Docker compose creates 2 containers: postgres & api.
<br> - ```postgres``` is the base container which is built up from ```postgres:14-alpine``` with all environment variables provided. ```Port: 5432:5432``` is to share this service outside of the postgres container. ```EXPOSE``` is to set its avalability to current container.
<br> - ```api``` is built from setting in Dockerfile in current folder with port 8080:8080. Environment is modified with command ```DB_SOURCE=postgresql://root:sec...``` to change the network of api to the network of postgres. This container will be built after postgres. Entrypoint is overwrtitten to be started in order.

# Demo API
### Testing api with postman & results:

The service is up and running:
![Screen Shot 2022-05-30 at 13 32 16](https://user-images.githubusercontent.com/106065029/170930852-760cc7b0-56fa-4d55-842a-049d1a3d59bd.png)

<br> GET: ```http://localhost:8080/multipleorder``` #Auto-generate sample data & get 5 sample records.

![Screen Shot 2022-05-30 at 13 26 06](https://user-images.githubusercontent.com/106065029/170930060-ec1a8a00-b57b-4848-ae92-220241ade09b.png)

<br> POST: ```http://localhost:8080/order``` #Create singular record {"user":string, "product":string, "amount":int}
![Screen Shot 2022-05-30 at 13 28 08](https://user-images.githubusercontent.com/106065029/170930332-a02f2be4-ff3c-481b-9cc0-c8ec5f3df5bc.png)

<br> GET: ```http://localhost:8080/order```

![Screen Shot 2022-05-30 at 13 20 04](https://user-images.githubusercontent.com/106065029/170930210-94c3edd6-fc01-4c13-a59a-6644540f04ec.png)

<br> Checking Postgres database:
<img width="1430" alt="Screen Shot 2022-05-30 at 13 29 07" src="https://user-images.githubusercontent.com/106065029/170930470-f3859ac7-1ff7-4d19-b024-d9ea9b824a13.png">

<br> Checking result generated by terminal:
![Screen Shot 2022-05-30 at 13 30 12](https://user-images.githubusercontent.com/106065029/170930599-64fb1352-b9d8-469f-ba1e-fbb97eadcba2.png)
