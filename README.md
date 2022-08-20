# Task API Gateway
### Installing the application
Clone repository
### Running the application
Run `docker-compose.yml`
### Testing and requests
#### The server handles requests of the type:
1) `localhost:8081/microservice/name` request to User Microservice and output information about the name of the microservice
2) `localhost:8081/user/profile?username=` request /auth to Microservice Authorization. If successful, request to User Microservice and output information about the user with a definite Username
