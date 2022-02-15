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
  "event_name": "BarEvent"
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
To implement the story, we are developing the system using golang.
The system will have two microservices, namely `fd-event-listener` and `fd-auth`.
Both the microservice will interact with the same database named `fd-db`. We will be using
postgresql for storing our data.

Here we are illustrating the high-level-diagram of the system

![Alt text](./doc-file/test.svg)
