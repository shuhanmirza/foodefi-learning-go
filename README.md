# foodefi-alpine

## Story: Event API

The FooDeFi company has a lot of scrapers that are scraping blockchains for Events. We want to design an API so that
these scrapers can submit their scraped Events and Admins can aggregate Events easily.

The Events scraped by scrapers has the following format

```
{
    "blockchain_id": <a number>,
    "block_number": <a number>,
    "event_name": <a string>,
    "fields": [
    {
        "type": <oneof number|string|boolean>,
        "name": <a string>,
        "value": <value according to type>
    },
    ... several fields of above format
    ]
}
```

Example 1:

```json
{
  "blockchain_id": 47,
  "block_number": 3916102,
  "event_name": "BarEvent",
  "fields": [
    {
      "type": "number",
      "name": "field_1",
      "value": 82
    },
    {
      "type": "string",
      "name": "field_2",
      "value": "demo"
    }
  ]
}
```

As a Backend Developer at FooDeFi, you are given the following requirements:

- The API should have two roles. Admins and Scrapers. 
- There should be an endpoint that allows only Scrapers to submit Events. 
- There should be endpoints to update / delete Events, both Admins and Scapers should be able to access this endpoint. 
- An endpoint should allow only Admins to list events. The endpoint should optionally allow to filter events by
  - blockchain_id 
  - event_name 
  - block_number range (should return Events that has block number in the range)


## Architecture
To implement the story, I am developing the system using golang.
The system will have two microservices, namely `fd-event-listener` and `fd-auth`.
Both the microservice will interact with the same database named `fd-db`. I will be using
postgresql for storing the data. To separate the concerns and loads, I chose to divide the system 
into microservices. 

Here I am illustrating the high-level-diagram of the system:

![Alt text](./doc-file/test.svg)

`fd-auth` will have the responsibility of authenticating the users and providing them with the jwt token. 
`fd-event-listener` will be responsible for managing the data of the events.

First, a user have to log in to the system through `fd-auth` and get jwt token.
Secondly, the user will use this token to interact with `fd-event-listener`. `fd-event-listener` will 
authorize the user based on the claim of the token. The token will contain various data along with 
username and the role.

### database schema
Here is the design of the database schema. I used dbdiagram.io to build the diagram.

![Alt text](./doc-file/db-digram.png)

For better view go to this [link](https://dbdiagram.io/d/620b828085022f4ee597fc93).

The database has four tables: (a) users, (b) blockchains, (c) events, and (d) event_fields. The schema above shows
relations between the tables. Event Fields were separated from Events for easing the connectivity with other tables
in the future and reduce data duplication.

## Getting Started
Have installed:
- golang
- docker and docker-compose
- golang-migrate
- sqlc
- postman

The project uses `gin` for implementing RESTful API and `sqlc` for database operations. 
Furthermore, `viper` was used for environment variable management.

Please find the postman collection in the `doc-file` folder

### Project Structure

    ├── README.md          <- you are currently reading this
    │── .gitignore         <- you know what this is :P
    ├── sqlc.yaml          <- go-sqlc configration
    ├── doc-file           <- contains the files that is necessary for documenting the project
    ├── docker             <- find the docker-compose files here
    │
    ├── db
    │   ├── migration         <- contains database migrations scripts
    │   ├── query             <- contains database sql scripts for go-sqlc
    │   └── foo-defi-init.sql <- initial script for creating all the tables 
    │
    └── foo-defi-go                   
        ├── main.go        <- main function for deploying the project
        ├── go.mod         <- go module file for describing module's properties, dependencies etc
        ├── app.env        <- environment variables for the project
        ├── util           <- common utilty package        
        │   ├── config.go     <- for loading configuration from the .env file
        │   ├── constants.go  <- global constants for the project
        │   ├── crypto.go     <- functions for cryptographic operations 
        │   ├── errors.go     <- global exceptions/error classes for the project
        │   ├── functions.go  <- utility functions for the project
        │   └── randoms.go    <- for producing random values
        │   
        ├── sqlc           <- for all database operations        
        │   ├── *.sql.go      <- sqlc generated code. contains database query functions
        │   ├── db.go         <- sqlc generated code. contains interface and class definitions
        │   ├── models.go     <- sqlc generated code. contains models for the tables
        │   └── store.go      <- for accessing all database classes. Also, contains db transactions
        │
        ├── fd-auth           
        │   ├── server.go     <- routing calls to handlers and server definitions
        │   └── login_api.go  <- handler function for serving login api for users
        │
        └── fd-event-listener           
            ├── server.go      <- routing calls to handlers and server definitions
            ├── submit_api.go  <- handler function for serving api for event submission 
            ├── update_api.go  <- handler function for serving api for updating event fields 
            ├── delete_api.go  <- handler function for serving api for event deletion 
            └── submit_api.go  <- handler function for serving api for retriving events 


## Note
This section will have the TODOs, improvement scopes, acknowledgment and confessions.

### TODOs
- JWT Token generation and authorization of tokens
- mitigating brute-force attack on login api
- minimize code duplication
- refactor the project for implementing a state-of-the-art design pattern for the project
- write tests to get full test coverage

### Confession and Improvement Scopes
This is the first web application that I implemented with golang. There's a lot of TODOs and improvement 
scope in this project.
- golang's error handling seemed quite unique for a java programmer like me. I could have avoided a lot of code 
duplication if I knew the idiomatic way of handling errors in golang development
- I could not structure the project into two microservices as I intended to. I need to learn how to manage 
multiple golang microservices in single repository. However, I leveraged go-routines to run two separate `gin`
services
- `sqlc` is love
- `viper` seemed confusing for me
- validation capability of `gin` was really confusing. I should have written custom validation classes
- golang's concept of classes, interfaces is pretty unique. Need to deep dive into this
- I could not learn and implement any design pattern for golang web development
- In real life, I practice FDD. This was my first time doing TDD
- I could not finish implementing JWT token authentication and authorization. You will see that I have kept
`shuhan_scraper` as recorder of all submitted event in `submit_api.go`. you must have a scraper named `shuhan_scraper`
in `users` table.

### Acknowledgement
- Mr. Quang Pham from TECH SCHOOL (https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)
- go.dev , dbdiagram.io, jenkins, sqlc.dev, w3schools.com, google.com,youtube.com, github.com


-----


## Utility Commands
So that, I do not have to search google for commands every time :p

```shell
docker exec -it postgres14 psql -U root
docker logs postgres14
```
```shell
brew install golang-migrate
migrate create -ext sql -dir db/migration -seq init_schema
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fd-db?sslmode=disable" -verbose up  
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fd-db?sslmode=disable" -verbose down
```
```shell
brew install sqlc
sqlc init
sqlc generate
```
```shell
go test -v -cover ./...
go run main.go
```
