# Glassnode Challenge
* Provide an API to get how much fees in the Ethereum network have been spent by plain ETH transfers between **EOA**.
* Use golang to build a restful api server to access data from PostgreSQL. 
* The service need to be started with the provided `docker-compose.yaml`.


## Solution
### Requirements
[Docker engine](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/) are needed for this service.

### Installation instructions / Run Instructions

I uploaded my backend image to docker hub.

Simply run this command in cli. The service will be ready after **1-2 minutes** for the setting up of the database.  
```bash
  docker-compose up
```
Stop the service by the following command.
```bash
  docker-compose down
```


### Testing steps
After the service is started, you can copy this link into any browsers or Postman to get the result.

HTTP Request Get
```html
  http://localhost:8080/eth/gas_hourly
```

### Result 
Where `t` is a unix timestamp of the hour, and `v` is the amount of fees being paid for transactions between **EOA** addresses within that hour in ETH units. It is in descending order.
```Json
  [{"t":1599519600,"v":16.27496391263491},
  {"t":1599516000,"v":13.956868558000194},
  {"t":1599512400,"v":17.011253750710875},
  {"t":1599508800,"v":19.009858770299253},
  {"t":1599505200,"v":19.52287467995422},
  {"t":1599501600,"v":22.009971273466867},
  {"t":1599498000,"v":25.38757770943181},
  {"t":1599494400,"v":34.39976542736484},
  {"t":1599490800,"v":30.78263919552725},
  {"t":1599487200,"v":37.934713446764775},
  {"t":1599483600,"v":34.36798373029837},
  {"t":1599480000,"v":28.445457754715306},
  {"t":1599476400,"v":26.965695348910778},
  {"t":1599472800,"v":28.731211524168447},
  {"t":1599469200,"v":30.734628412396855},
  {"t":1599465600,"v":28.97972350167535},
  {"t":1599462000,"v":34.50883497287487},
  {"t":1599458400,"v":32.02156940081887},
  {"t":1599454800,"v":24.498179986606804},
  {"t":1599451200,"v":26.799904503862056},
  {"t":1599447600,"v":28.011215510069206},
  {"t":1599444000,"v":33.47291557101054},
  {"t":1599440400,"v":25.10686991549367},
  {"t":1599436800,"v":16.882940469082822}]
```

### Query Statement
```sql
  SELECT 
    date_trunc('hour', t.block_time) as block_time_hour, 
    sum(t.gas_used * t.gas_price) / 10 ^ 18 as eth
  FROM 
    public.transactions t
  WHERE 
    t.to != '0x0000000000000000000000000000000000000000'
    AND t.value > 0
    AND not exists (select 1 from public.contracts c WHERE t.to = c.address)
  GROUP BY block_time_hour
  ORDER BY block_time_hour DESC;
```

For this challenge, I am looking for all the plain ETH tranfer between **EOA** only. Here is the explaination of the WHERE clause condition.

1. Transaction of contract creation is from **EOA** to a special address `0x0000000000000000000000000000000000000000`. The condition should be `t.to != '0x0000000000000000000000000000000000000000'`
2. The `value` variable in the `trasactions` table must be greater than 0, since any transaction with value 0 is not transfering any ether between the `from` and `to` address. The condition should be `t.value > 0`
3. For any ether transfers from **EOA** to contract address, it will be recorded into the block transaction. And any ether transfer from contract address to **EOA**/contract address, the record will be stored in the `message call`. Contract address will only exist in the `to` in the transactions. The condition should be `not exists (select 1 from public.contracts c WHERE t.to = c.address)`

## Approaches and Tradeoffs

### Technical choices
* GoLang 1.14
* github.com/gofiber/fiber v1.14.5
* github.com/sirupsen/logrus v1.8.1
* github.com/lib/pq v1.10.0

#### 1. Golang
I chose GoLang 1.14 since I had experience with it before. Since the code challenge is deigned for 3-4 hours, I don't want to deal with the problem of different version of GoLang, i.e. the version of GoLang in my compiler.

#### 2. Web framework
Fiber is my choice of web framework in this tasks. Based on their [performance tests](https://docs.gofiber.io/extra/benchmarks), fiber is the the fastest HTTP engine for Go.

#### 3. Logging
github.com/sirupsen/logrus provide different level of logging and good representation.

#### 4. DB connection
github.com/lib/pq provide the connection setting to connect the required PostgreSQL database.

### System Design
Backend application design is based on MVC. My code is devided into 5 packages which are `main`, `route`, `model`, `handler`, `databaes`. And I made test cases for the main function.

* `main` is the project main starting point. 
* `route` is setting all the rest api route. It routes all the request that match a specific path.
* `model` is the package of the data transfer object. It is for transferring data between software application subsystems. 
* `handler` is for the api endpoint response. It controls the requets response of every indivilial apis.
* `database` includes all the functions needed to comunnicate with database. The establishment of connection is in `connection.go`.  The query statement is in `eth.go`. 

Since this task only requires to build single GET api without extra function, I didn't include the `middeware`, `businesss layer` package  and building `interface` in the code. 

* For `middleware`, it will intercept the api request before it reaches the `handler` layer, i.e. the public api can limit repeated requests for the same IP address. 

* For `businesss layer`, I will place it between the `handler` and `database` layer. It is used for more complex logic, i.e. calling mutliple functoin in `database` layer. 

* For `interface`, it provides the implementations of the same object or function, i.e. building the interface of the `database` layer. It can provide an interface to connect to different database - postgreSQL, mySQL, since 2 different database requre different library and config.

## Summary
The greatest challenge for me was to understand `Plain ETH trasfer between EOAs`. I spent most of the time researching the transfer between the contract address and EOA. In the beginning of the task, I used `left octer join` twice by joining `to` and `from` to the contract address. It slows down the query statement and I wanted to optimize it. Finally, I found out I only have to join `to` to contact address. 

Thank you for designing this interesting code challenge. I am really looking forward to discuss my solution. Since I am still in a learning process using Golang, I would appreciate your feedback on my coding skills.
