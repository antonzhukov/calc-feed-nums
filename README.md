# Calculate feed numbers
This is the test task I made for one of the companies
##Run application
Run 2 webservers:
1. ```go run main.go```
2. ```go run feedserver/main.go```

Test application:
```
http://localhost:8080/numbers?u=http://localhost:8090/primes&u=http://localhost:8090/fibo
```
##Description
The solution is built upon following architecture: 
service listens on 8080 port, it has a router with all 
the possible routes and corresponding services to 
handle the requests. 
All json structures are described as Data Transfer Objects.
Requests to feed services are made in concurrency mode

##Problems, ambiguities and decisions
- It was not mentioned how to handle unavailable feeds so these incidents are just written to log
- I didn't bother to use regexp to precisely check if url is valid, because service needs to respond as quickly as possible
- Service max execution time was claimed 500ms so feed urls timeout set to 450ms

##TODO
- tests

by Anton Zhukov