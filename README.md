### Architecture

This project is to experiment to develop distributed application with Go. The primary object of the project is to use 
use API gateway to perform all the different plumbing such as - routing, security, JSON processing, etc.

The project will use few cool technologies to solve distributed issues such as async communicataion, service resolution,
caching and many more.

The project uses [KrakenD project](https://github.com/krakend/krakend-ce)

![architecture](docs/architecture.png)


### Installing and Running

* Register in https://comicvine.gamespot.com/api/ to register and get an API key.

* Run rabbitmq

```
docker run -p 5672:5672 -p 15672:15672 datamanipulation/rabbitmq:3.11-management-alpine
```

* Run krakend-ce from inside `comics` directory:

```
krakend run -d -c gateway/comic.json
```

* Run `main.go`

```
go run main.go -token=<comicvine_token>
```

The parameter `-token` contains the Comicview API key. The application will run in a loop 3 times, each time will obtain 
10 records, with the records pushed to the RabbitMQ queue - `producer-q-exchange`.


### Accessing endpoint

The endpoint `/comic` is exposed by the KrakenD gateway, which can be accessed using cURL:

```
curl http://localhost:8080/comic |  jq . 
```
