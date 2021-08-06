# NeoPaginal

- This shows you how to crawls pages steps for 3 times, store and projection of detail about pages with microservices architecture.

## Architecture

![diagram-single-pod](https://raw.githubusercontent.com/oktydag/NeoPaginal/main/contents/architecture.png)

## Project Descriptions

#### **neopaginal.command**
This crawls a given page requested times( Default 3 ) in succession and stores maintained data to MongoDB.
Keys; 
- Golang
- Golang Naming Convension Best Practices (https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go/25162021)
- Microservices, DDD, BDD
- TDD, Unit Testing, Mocking, Testify Assertion
- MongoDB ( Write Database )
- Builder (To Create By Prepared Rules), Modular Design Pattern
- CQRS ( 3.Level of CQRS : separated storage - https://levelup.gitconnected.com/3-cqrs-architectures-that-every-software-architect-should-know-a7f69aae8b6c)
- Docker


#### **neopaginal.passanger**
This reads data from MongoDB, enriches them and write to projection database that is Elasticsearch as bulk. 
- Golang
- Golang Naming Convension Best Practices (https://stackoverflow.com/questions/25161774/what-are-conventions-for-filenames-in-go/25162021)
- Microservices, BDD
- MongoDB ( Write Database ) , ElasticSearch ( Read Database )
- CQRS ( 3.Level of CQRS : separated storage - https://levelup.gitconnected.com/3-cqrs-architectures-that-every-software-architect-should-know-a7f69aae8b6c)
- Docker

#### **neopaginal.query.api**
This is restful web api that datasource is Elasticsearch 
- Nodejs
- Restful Web Api with Express
- Restful Web Api, Naming Convension, Api Versioning Best Practices ( https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design )
- ElasticSearch ( Read Database )
- CQRS ( 2.Level of CQRS : Two-database - https://levelup.gitconnected.com/3-cqrs-architectures-that-every-software-architect-should-know-a7f69aae8b6c)
- Swagger Documentation ( address : https://0.0.0.0:5002/api-docs )
- Docker

### Run microservices

In directory that contains docker-compose.yaml;

<pre> $ docker-compose up --build
</pre>

# Best Practice Notes
- neopaginal.command can process the crawl and create a Outbox Message template to write OutboxMessage collection to to protect messages in any case (Outbox Microservice Pattern : https://microservices.io/patterns/data/transactional-outbox.html)

- Outbox.Publisher app could be created to read message from OutboxMessage collection and messages of events have published to Queue with any messsage broker such as Kafka, RabbitMQ, WebSphereMQ etc.. ( Event Sourcing : https://martinfowler.com/eaaDev/EventSourcing.html#:~:text=The%20fundamental%20idea%20of%20Event,as%20the%20application%20state%20itself. )

- neopaginal.passanger could be a consumer that subscribe queue and write messages MongoDB and Elasticsearch to seperate read and write databases ( Event Sourcing with CQRS : https://microservices.io/patterns/data/event-sourcing.html )
