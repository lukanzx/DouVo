
#### DouVo
This is a simple short video service backend based on HTTP and RPC protocols, distributed architecture


#### Technology stack
 - Kitex
 - Hertz
 - etcd
 - MySQL
 - Redis
 - Jaeger
 - Prometheus
 - Grafana
 - Elasticsearch
 - Kibana


#### Implemented functional modules
 - Like function module
 - Chat function module
 - Video function module
 - Comment function module
 - Focus on functional modules


#### characteristic
 - ##### Modularity and Independence: 
    Microservices allow teams to break down applications into small modules that can be independently developed, deployed, and extended.
 - ##### Easy to expand: 
    Individual microservices can be independently expanded as needed, without the need to expand the entire application. This makes microservices architecture very suitable for handling uneven loads between different services.
 - ##### Fault tolerance: 
    The failure of a single microservice usually does not affect the entire application. This means that the system can be more fault-tolerant, thereby improving availability.


#### Architecture
![architecture](./docs/img/AAarchitecture.png)


#### Project Structure


##### whole
```
.

├── cmd                   # Microservices
├── config                # allocation
├── docs
├── go.mod
├── go.sum
├── kitex_gen
├── pkg
│   ├── constants         # Store some constants
│   ├── errno             # Customization error
│   ├── middleware        # Common middleware
│   ├── tracer            # Link tracking
│   └── utils             # Useful features
└── test
```

##### Gateway/API module
```
.
├── Makefile
├── biz
│   ├── handler     # Request processing, packaging, and returning data
│   ├── middleware
│   ├── model
│   ├── pack        # pack
│   ├── router      # route
│   └── rpc         # Send RPC request
├── build.sh
├── main.go
├── output          # Binary files
├── router.go
├── router_gen.go
└── script
```

##### Microservices
```
.
├── Makefile        # Useful commands
├── build.sh        # Building scripts
├── dal
│   ├── cache       # redis
│   ├── db          # MySQL
│   └── mq          # RabbitMQ
├── handler.go
├── kitex_info.yaml
├── main.go
├── output          # Built binary files
├── pack            # Pack and return data
├── rpc             # Send requests to other RPC services
├── script
├── coverage        # Coverage test results (some not available)
└── service
```
