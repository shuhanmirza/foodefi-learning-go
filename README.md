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
postgresql for storing the data.

Here I am illustrating the high-level-diagram of the system

![Alt text](./doc-file/test.svg)

### database schema
Here is the design of the database schema. I used dbdiagram.io to build the diagram.

![Alt text](./doc-file/db-digram.png)

For better view go to this [link](https://dbdiagram.io/d/620b828085022f4ee597fc93).

## Getting Started
Have installed:
- golang
- docker and docker-compose
- golang-migrate
- sqlc

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
```
