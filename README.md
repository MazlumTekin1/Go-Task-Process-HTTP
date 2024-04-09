# Go Task Process Net/Http

## Uyarı: Şu anda Swagger ile ilgili ufak bir sıkıntı var. Yakında çözülecek. Eğer bu süre zarfı içinde bakıyorsanız internal/sql klasöründeki Go-Task-Project-Http-Postman-Collection.json collection'u import edebilirsiniz.
## Warning: There is a slight issue with Swagger right now. It will be resolved soon

## Getting Started

- I built the project on Hexagonal Architecture. 
- Inside the Internal folder is the API folder with the request ends,
- Handler folder that receives the incoming request,
- Service folder with business rules,
- Repository folder for database operations,
- There is a Domain folder for the Model structures of Request and Response operations.

- In this project, all CRUD operations of Users and Tasks are prepared and a system that calculates the number of Tasks that Users can get per week is designed.
- You need to pass JWT authentication from /login in order to send requests to the request ends. There is a JWT token query on all existing request ends except login. If there is no token or if it is not passed, the request will not be made.

- RateLimit function has been added to all request ends. A middleware business rule has been created that can make a maximum of 10 requests in 1 minute.

- With the Bearer Token you will receive after login, you can add/delete/update/list Users (All Users and Id-based User retrieval).

- In the same way, you can do Add/Delete/Update/List (All Task and Id-based Task fetching) operations for Task. 

- A Linear Programming function has been designed that can distribute all tasks in the shortest time by taking into account the levels of Users, the difficulty levels of Tasks and Task durations, so that it can only do tasks at its own level and lower levels.

- Swagger implementation was prepared for API documentation.
- Promethius tool was implemented for monitoring. Metrics and count feature can be monitored.

- Created unit tests for User and Task

### Prerequisites

What things you need to install the software and how to install them.

- Go version (after 1.20)
- Docker

### Installing

A step by step series of examples that tell you how to get a development environment running.

1. Clone the repository
2. Run `go mod download` to download dependencies
3. For database connection, you need to replace the /internal/config.json file with your own PostgreSQL database information.
4. You need to run DDL commands in the /internal/sql folder. Or you can import the dump file below
5. go to root project directory
6. cd cmd
7. go run main.go


## Running the tests

Explain how to run the automated tests for this system.

- For unit tests, you can run `go test ./...`

## Deployment

```bash
    docker build -t go-task-process-http .
```

```bash
    docker compose up
```


## Authors

- Mazlum Tekin - Initial work - [MazlumTekin](https://github.com/MazlumTekin1)
